package caesar

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func caesarEncrypt(text string, shift int) string {
	shift = shift % 26
	encryptedText := make([]rune, len(text))
	for i, char := range text {
		if char >= 'a' && char <= 'z' {
			encryptedText[i] = 'a' + (char-'a'+rune(shift))%26
		} else if char >= 'A' && char <= 'Z' {
			encryptedText[i] = 'A' + (char-'A'+rune(shift))%26
		} else {
			encryptedText[i] = char
		}
	}
	return string(encryptedText)
}

func caesarDecrypt(text string, shift int) string {
	return caesarEncrypt(text, -shift)
}
