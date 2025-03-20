package threedes

import (
	"bufio"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/lukerten/cypher/src/config"

	"github.com/spf13/cobra"
)

func New3DESCommand() *cobra.Command {
	threedesCmd := &cobra.Command{
		Use:   "3des",
		Short: "3DES encryption and decryption commands",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt text using 3DES",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig("config/config.yml")
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			key := []byte(config.TripleDES.Key)

			text := getInputText(args)
			encryptedText, err := tripleDESEncrypt(text, key)
			if err != nil {
				fmt.Println("Error encrypting text:", err)
				os.Exit(1)
			}
			fmt.Println(encryptedText)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt text using 3DES",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig("config/config.yml")
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			key := []byte(config.TripleDES.Key)

			text := getInputText(args)
			decryptedText, err := tripleDESDecrypt(text, key)
			if err != nil {
				fmt.Println("Error decrypting text:", err)
				os.Exit(1)
			}
			fmt.Println(decryptedText)
		},
	}

	threedesCmd.AddCommand(encryptCmd, decryptCmd)
	return threedesCmd
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

func tripleDESEncrypt(text string, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create 3DES cipher: %w", err)
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, key[:block.BlockSize()])
	stream.XORKeyStream(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func tripleDESDecrypt(text string, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create 3DES cipher: %w", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 text: %w", err)
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, key[:block.BlockSize()])
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
