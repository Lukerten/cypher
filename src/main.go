package main

import (
	"fmt"
	"os"

	threedes "github.com/lukerten/cypher/src/3des"
	"github.com/lukerten/cypher/src/atbash"
	"github.com/lukerten/cypher/src/base64"
	"github.com/lukerten/cypher/src/caesar"
	"github.com/lukerten/cypher/src/rot13"
	"github.com/lukerten/cypher/src/rsa"
	"github.com/lukerten/cypher/src/substitution"

	"github.com/spf13/cobra"
)

func main() {
	var configPath string

	rootCmd := &cobra.Command{
		Use:   "cypher",
		Short: "Cypher is a simple terminal based command",
		Long:  `Cypher is a simple terminal based command for encryption and decryption.`,
	}

	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config/config.yml", "path to the config file")

	base64Cmd := base64.NewBase64Command()
	substitutionCmd := substitution.NewSubstitutionCommand()
	rsaCmd := rsa.NewRSACommand()
	threedesCmd := threedes.New3DESCommand()
	caesarCmd := caesar.NewCaesarCommand()
	atbashCmd := atbash.NewAtbashCommand()
	rot13Cmd := rot13.NewROT13Command()

	rootCmd.AddCommand(base64Cmd, substitutionCmd, rsaCmd, threedesCmd, caesarCmd, atbashCmd, rot13Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
