package cmd

import (
	"git.jcu.edu.au/cft/cfds/crypt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Given a seed and a password, will decrpyt to STDOUT",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		decrypted, err := crypt.Decrypt(args[0], seed)
		if err != nil {
			panic(err)
		}
		println(decrypted)
	},
}
