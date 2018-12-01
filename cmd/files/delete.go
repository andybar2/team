package files

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var deleteParams struct {
	Stage string
	Path  string
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file",
	RunE:  runDeleteCmd,
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteParams.Stage, "stage", "s", "", "Stage name")
	deleteCmd.Flags().StringVarP(&deleteParams.Path, "path", "p", "", "File path")

	FilesCmd.AddCommand(deleteCmd)
}

func runDeleteCmd(cmd *cobra.Command, args []string) error {
	if deleteParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if deleteParams.Path == "" {
		return errors.New("invalid file path")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.FileDelete(deleteParams.Stage, deleteParams.Path)
}
