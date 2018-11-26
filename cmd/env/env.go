package env

import (
	"github.com/spf13/cobra"
)

// EnvCmd is the sub-command to manage environment veriables
var EnvCmd = &cobra.Command{
	Use:   "env [command]",
	Short: "Manage environment variables",
}
