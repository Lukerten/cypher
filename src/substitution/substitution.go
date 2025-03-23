package substitution

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
