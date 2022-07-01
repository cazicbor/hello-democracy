package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello-democracy",
	Short: "hello-democracy",
	Long:  `Implementation of various voting methods`,
}

func Execute() {
	rootCmd.AddCommand(
		majorityCmd,
		approvalCmd,
		condorcetCmd,
		copelandCmd,
		fullCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
