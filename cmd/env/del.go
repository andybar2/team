package env

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var delParams struct {
	Stage string
	Name  string
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an environment variable",
	RunE:  runDelCmd,
}

func init() {
	delCmd.Flags().StringVarP(&delParams.Stage, "stage", "s", "", "Stage name")
	delCmd.Flags().StringVarP(&delParams.Name, "name", "n", "", "Variable name")

	EnvCmd.AddCommand(delCmd)
}

func runDelCmd(cmd *cobra.Command, args []string) error {
	if delParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if delParams.Name == "" {
		return errors.New("invalid variable name")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.EnvDelete(delParams.Stage, delParams.Name)
}
