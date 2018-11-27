package cmd

import (
	"github.com/andybar2/team/cmd/env"
	"github.com/andybar2/team/cmd/file"
	"github.com/spf13/cobra"
)

// RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "team [command]",
	Short: "Store project configuration remotely and easily share it with your team",
}

// Execute runs the root command
func Execute() error {
	RootCmd.AddCommand(env.EnvCmd)
	RootCmd.AddCommand(file.FileCmd)

	return RootCmd.Execute()
}
