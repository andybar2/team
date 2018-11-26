package env

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var setParams struct {
	Stage string
	Name  string
	Value string
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an environment variable",
	RunE:  runSetCmd,
}

func init() {
	setCmd.Flags().StringVarP(&setParams.Stage, "stage", "s", "", "Stage name")
	setCmd.Flags().StringVarP(&setParams.Name, "name", "n", "", "Variable name")
	setCmd.Flags().StringVarP(&setParams.Value, "value", "v", "", "Variable value")

	EnvCmd.AddCommand(setCmd)
}

func runSetCmd(cmd *cobra.Command, args []string) error {
	if setParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if setParams.Name == "" {
		return errors.New("invalid variable name")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.EnvSet(setParams.Stage, setParams.Name, setParams.Value)
}
