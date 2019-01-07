package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the currently installed version of Janus",
	Long:  `Janus version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Janus Version %s", version)
	},
}