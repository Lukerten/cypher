package rsa

import (
	"fmt"
	"os"

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
