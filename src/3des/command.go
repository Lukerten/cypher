package threedes

import (
	"crypto/des"
	"fmt"
	"os"

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
			iv := []byte(config.TripleDES.IV)
			if len(key) != 24 {
				fmt.Println("Error: Key length must be 24 bytes")
				os.Exit(1)
			}
			if len(iv) != des.BlockSize {
				fmt.Println("Error: IV length must be 8 bytes")
				os.Exit(1)
			}

			text := getInputText(args)
			encryptedText, err := tripleDESEncrypt(text, key, iv)
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
			iv := []byte(config.TripleDES.IV)
			if len(key) != 24 {
				fmt.Println("Error: Key length must be 24 bytes")
				os.Exit(1)
			}
			if len(iv) != des.BlockSize {
				fmt.Println("Error: IV length must be 8 bytes")
				os.Exit(1)
			}

			text := getInputText(args)
			decryptedText, err := tripleDESDecrypt(text, key, iv)
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
