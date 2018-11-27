package file

import (
	"github.com/spf13/cobra"
)

// FileCmd is the sub-command to manage configuration files
var FileCmd = &cobra.Command{
	Use:   "file [command]",
	Short: "Manage configuration files",
}
