package cmd

import (
	"git.jcu.edu.au/cft/cfds/crypt"
	"github.com/ansel1/merry"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Given a seed and a password, will decrpyt to STDOUT",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if seed == "" {
			return merry.New("Requires seed parameter")
		}
		decrypted, err := crypt.Decrypt(args[0], seed)
		if err != nil {
			return err
		}
		println(decrypted)
		return nil
	},
}
