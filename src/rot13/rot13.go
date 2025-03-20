package rot13

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func rot13(text string) string {
	rotated := make([]rune, len(text))
	for i, char := range text {
		if char >= 'a' && char <= 'z' {
			rotated[i] = 'a' + (char-'a'+13)%26
		} else if char >= 'A' && char <= 'Z' {
			rotated[i] = 'A' + (char-'A'+13)%26
		} else {
			rotated[i] = char
		}
	}
	return string(rotated)
}
