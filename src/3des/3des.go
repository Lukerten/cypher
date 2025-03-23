package threedes

import (
	"bufio"
	"bytes"
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
	padding := block.BlockSize() - len(plaintext)%block.BlockSize()
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

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
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove padding
	padding := int(plaintext[len(plaintext)-1])
	plaintext = plaintext[:len(plaintext)-padding]

	return string(plaintext), nil
}
