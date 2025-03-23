package rot13

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
