package substitution

import (
	"fmt"
	"os"

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
