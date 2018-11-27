package file

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var uploadParams struct {
	Stage string
	Name  string
	Path  string
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file",
	RunE:  runUploadCmd,
}

func init() {
	uploadCmd.Flags().StringVarP(&uploadParams.Stage, "stage", "s", "", "Stage name")
	uploadCmd.Flags().StringVarP(&uploadParams.Name, "name", "n", "", "File name")
	uploadCmd.Flags().StringVarP(&uploadParams.Path, "path", "p", "", "File path")

	FileCmd.AddCommand(uploadCmd)
}

func runUploadCmd(cmd *cobra.Command, args []string) error {
	if uploadParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if uploadParams.Name == "" {
		return errors.New("invalid file name")
	}

	if uploadParams.Path == "" {
		return errors.New("invalid file path")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.FileUpload(uploadParams.Stage, uploadParams.Name, uploadParams.Path)
}
