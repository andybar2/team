package files

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var downloadParams struct {
	Stage string
	Path  string
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a file",
	RunE:  runDownloadCmd,
}

func init() {
	downloadCmd.Flags().StringVarP(&downloadParams.Stage, "stage", "s", "", "Stage name")
	downloadCmd.Flags().StringVarP(&downloadParams.Path, "path", "p", "", "File path")

	FilesCmd.AddCommand(downloadCmd)
}

func runDownloadCmd(cmd *cobra.Command, args []string) error {
	if downloadParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if downloadParams.Path == "" {
		return errors.New("invalid file path")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.FileDownload(downloadParams.Stage, downloadParams.Path)
}
