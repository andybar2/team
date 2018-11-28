package files

import (
	"github.com/spf13/cobra"
)

// FilesCmd is the sub-command to manage configuration files
var FilesCmd = &cobra.Command{
	Use:   "files [command]",
	Short: "Manage configuration files",
}
