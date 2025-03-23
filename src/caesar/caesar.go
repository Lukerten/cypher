package caesar

import (
	"bufio"
	"os"
	"strings"
)

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
