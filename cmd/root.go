package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "skim",
		Short: "skim is a CLI tool that extracts a list of container images from Kubernetes resources",
	}
	return rootCmd
}

func Execute() {
	rootCmd := newRootCmd()
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.SilenceUsage = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
