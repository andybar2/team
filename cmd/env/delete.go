package env

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var deleteParams struct {
	Stage string
	Name  string
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an environment variable",
	RunE:  runDeleteCmd,
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteParams.Stage, "stage", "s", "", "Stage name")
	deleteCmd.Flags().StringVarP(&deleteParams.Name, "name", "n", "", "Variable name")

	EnvCmd.AddCommand(deleteCmd)
}

func runDeleteCmd(cmd *cobra.Command, args []string) error {
	if deleteParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if deleteParams.Name == "" {
		return errors.New("invalid variable name")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.EnvDelete(deleteParams.Stage, deleteParams.Name)
}
