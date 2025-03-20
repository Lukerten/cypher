package atbash

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func atbash(text string) string {
	translated := make([]rune, len(text))
	for i, char := range text {
		if char >= 'a' && char <= 'z' {
			translated[i] = 'z' - (char - 'a')
		} else if char >= 'A' && char <= 'Z' {
			translated[i] = 'Z' - (char - 'A')
		} else {
			translated[i] = char
		}
	}
	return string(translated)
}
