package substitution

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lukerten/cypher/src/config"
	"github.com/spf13/cobra"
)

func NewSubstitutionCommand() *cobra.Command {
	substitutionCmd := &cobra.Command{
		Use:   "substitution",
		Short: "Substitution cipher commands",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt text using substitution cipher",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig("config/config.yml")
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			keys, err := LoadSubstitutionKeys(config.Substitution.Keys)
			if err != nil {
				fmt.Println("Error loading substitution keys:", err)
				os.Exit(1)
			}

			text := getInputText(args)
			encryptedText := substitutionEncrypt(text, keys)
			fmt.Println(encryptedText)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt text using substitution cipher",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig("config/config.yml")
			if err != nil {
				fmt.Println("Error loading config:", err)
				os.Exit(1)
			}

			keys, err := LoadSubstitutionKeys(config.Substitution.Keys)
			if err != nil {
				fmt.Println("Error loading substitution keys:", err)
				os.Exit(1)
			}

			text := getInputText(args)
			decryptedText := substitutionDecrypt(text, keys)
			fmt.Println(decryptedText)
		},
	}

	substitutionCmd.AddCommand(encryptCmd, decryptCmd)
	return substitutionCmd
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

func substitutionEncrypt(text string, keys []SubstitutionKey) string {
	encryptedText := text
	for _, key := range keys {
		encryptedText = strings.ReplaceAll(encryptedText, key.Char, key.Value)
	}
	return encryptedText
}

func substitutionDecrypt(text string, keys []SubstitutionKey) string {
	decryptedText := text
	for _, key := range keys {
		decryptedText = strings.ReplaceAll(decryptedText, key.Value, key.Char)
	}
	return decryptedText
}
