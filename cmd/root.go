package cmd

import (
	"github.com/spf13/cobra"
)

const configFile = "team-env.json"

var rootCmd = &cobra.Command{
	Use:   "team-env [command]",
	Short: "Store environment variables remotely and share them with your team",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
