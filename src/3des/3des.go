package threedes

import (
	"bufio"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
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

func tripleDESEncrypt(text string, key, iv []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create 3DES cipher: %w", err)
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func tripleDESDecrypt(text string, key, iv []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create 3DES cipher: %w", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 text: %w", err)
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
