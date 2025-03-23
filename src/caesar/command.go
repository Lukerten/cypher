package caesar

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func NewCaesarCommand() *cobra.Command {
	caesarCmd := &cobra.Command{
		Use:   "caesar",
		Short: "Caesar cipher commands",
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt [shift]",
		Short: "Encrypt text using Caesar cipher",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shift, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid shift value:", err)
				os.Exit(1)
			}

			text := getInputText(args[1:])
			encryptedText := caesarEncrypt(text, shift)
			fmt.Println(encryptedText)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt [shift]",
		Short: "Decrypt text using Caesar cipher",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shift, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid shift value:", err)
				os.Exit(1)
			}

			text := getInputText(args[1:])
			decryptedText := caesarDecrypt(text, shift)
			fmt.Println(decryptedText)
		},
	}

	caesarCmd.AddCommand(encryptCmd, decryptCmd)
	return caesarCmd
}
