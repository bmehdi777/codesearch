package cmd

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "codesearch",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(newStartCmd())

	return &rootCmd
}
