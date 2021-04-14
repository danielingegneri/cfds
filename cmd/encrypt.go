package cmd

import (
	"git.jcu.edu.au/cft/cfds/crypt"
	"github.com/ansel1/merry"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Given a seed and a password, will encrypt to STDOUT, encoded as base64",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if seed == "" {
			return merry.New("Requires seed parameter")
		}
		encrypted, err := crypt.Encrypt(args[0], seed)
		if err != nil {
			return err
		}
		println(encrypted)
		return nil
	},
}
