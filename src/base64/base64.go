package base64

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func NewBase64Command() *cobra.Command {
	base64Cmd := &cobra.Command{
		Use:   "base64",
		Short: "Base64 encoding and decoding commands",
	}

	encodeCmd := &cobra.Command{
		Use:   "encode",
		Short: "Encode text to base64",
		Run: func(cmd *cobra.Command, args []string) {
			text := getInputText(args)
			encodedText := base64.StdEncoding.EncodeToString([]byte(text))
			fmt.Println(encodedText)
		},
	}

	decodeCmd := &cobra.Command{
		Use:   "decode",
		Short: "Decode base64 to text",
		Run: func(cmd *cobra.Command, args []string) {
			text := getInputText(args)
			decodedBytes, err := base64.StdEncoding.DecodeString(text)
			if err != nil {
				fmt.Println("Error decoding base64:", err)
				os.Exit(1)
			}
			fmt.Println(string(decodedBytes))
		},
	}

	base64Cmd.AddCommand(encodeCmd, decodeCmd)
	return base64Cmd
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
