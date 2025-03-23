package atbash

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewAtbashCommand() *cobra.Command {
	atbashCmd := &cobra.Command{
		Use:   "atbash",
		Short: "Atbash cipher commands",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt text using Atbash cipher",
		Run: func(cmd *cobra.Command, args []string) {
			text := getInputText(args)
			encryptedText := atbash(text)
			fmt.Println(encryptedText)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt text using Atbash cipher",
		Run: func(cmd *cobra.Command, args []string) {
			text := getInputText(args)
			decryptedText := atbash(text)
			fmt.Println(decryptedText)
		},
	}

	atbashCmd.AddCommand(encryptCmd, decryptCmd)
	return atbashCmd
}
