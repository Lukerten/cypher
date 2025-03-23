package atbash

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
