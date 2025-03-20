package rsa

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/lukerten/cypher/src/config"

	"github.com/spf13/cobra"
)

func NewRSACommand() *cobra.Command {
	rsaCmd := &cobra.Command{
		Use:   "rsa",
		Short: "RSA encryption and decryption commands",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt text using RSA",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig("config/config.yml")
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			publicKey, err := loadRSAPublicKey(config.RSA.PublicKeyPath)
			if err != nil {
				fmt.Println("Error loading public key:", err)
				os.Exit(1)
			}

			text := getInputText(args)
			encryptedText, err := rsaEncrypt(text, publicKey)
			if err != nil {
				fmt.Println("Error encrypting text:", err)
				os.Exit(1)
			}
			fmt.Println(encryptedText)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt text using RSA",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig("config/config.yml")
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			privateKey, err := loadRSAPrivateKey(config.RSA.PrivateKeyPath)
			if err != nil {
				fmt.Println("Error loading private key:", err)
				os.Exit(1)
			}

			text := getInputText(args)
			decryptedText, err := rsaDecrypt(text, privateKey)
			if err != nil {
				fmt.Println("Error decrypting text:", err)
				os.Exit(1)
			}
			fmt.Println(decryptedText)
		},
	}

	rsaCmd.AddCommand(encryptCmd, decryptCmd)
	return rsaCmd
}

func getInputText(args []string) string {
	if len(args) > 0 {
		return strings.Join(args, " ")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var inputText strings.Builder
	for scanner.Scan() {
		inputText.WriteString(scanner.Text())
	}
	return inputText.String()
}

func loadRSAPublicKey(path string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return pub, nil
}

func loadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return priv, nil
}

func rsaEncrypt(text string, pub *rsa.PublicKey) (string, error) {
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(text))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt text: %w", err)
	}
	return string(encryptedBytes), nil
}

func rsaDecrypt(text string, priv *rsa.PrivateKey) (string, error) {
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(text))
	if err != nil {
		return "", fmt.Errorf("failed to decrypt text: %w", err)
	}
	return string(decryptedBytes), nil
}
