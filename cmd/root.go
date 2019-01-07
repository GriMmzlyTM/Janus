package cmd

import (
	"fmt"
	cobra "github.com/spf13/cobra"
	"os")

var rootCmd = &cobra.Command{
	Use:   "Janus",
	Short: "Janus SQL helper",
	Long: `The Janus SQL helper CLI app designed primarily for legacy applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}