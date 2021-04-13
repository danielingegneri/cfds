package cmd

import "github.com/spf13/cobra"

var seed string

var rootCmd = &cobra.Command{
	Use:   "cfds",
	Short: "CFDS can create Coldfusion datasource files and deal with encrypting and decrypting passwords",
}

func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&seed, "seed", "s", "", "0123456789ABCDEF")
	return rootCmd.Execute()
}
