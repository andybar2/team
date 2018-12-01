package files

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var listParams struct {
	Stage string
	Path  string
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files",
	RunE:  runListCmd,
}

func init() {
	listCmd.Flags().StringVarP(&listParams.Stage, "stage", "s", "", "Stage name")

	FilesCmd.AddCommand(listCmd)
}

func runListCmd(cmd *cobra.Command, args []string) error {
	if listParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.FileList(listParams.Stage)
}
