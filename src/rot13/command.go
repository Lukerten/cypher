package rot13

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewROT13Command() *cobra.Command {
	rot13Cmd := &cobra.Command{
		Use:   "rot13",
		Short: "ROT13 cipher commands",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt text using ROT13 cipher",
		Run: func(cmd *cobra.Command, args []string) {
			text := getInputText(args)
			encryptedText := rot13(text)
			fmt.Println(encryptedText)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt text using ROT13 cipher",
		Run: func(cmd *cobra.Command, args []string) {
			text := getInputText(args)
			decryptedText := rot13(text)
			fmt.Println(decryptedText)
		},
	}

	rot13Cmd.AddCommand(encryptCmd, decryptCmd)
	return rot13Cmd
}
