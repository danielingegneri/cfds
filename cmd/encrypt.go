package cmd

import (
	"git.jcu.edu.au/cft/cfds/crypt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Given a seed and a password, will encrypt to STDOUT, encoded as base64",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		encrypted, err := crypt.Encrypt(args[0], seed)
		if err != nil {
			panic(err)
		}
		println(encrypted)
	},
}
