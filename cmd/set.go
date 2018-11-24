package cmd

import (
	"errors"

	"github.com/andybar2/team-env/store"
	"github.com/spf13/cobra"
)

var serParams struct {
	Environment string
	Variable    string
	Value       string
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an environment variable",
	RunE:  runSetCmd,
}

func init() {
	setCmd.Flags().StringVar(&serParams.Environment, "env", "", "Environment name")
	setCmd.Flags().StringVar(&serParams.Variable, "var", "", "Variable name")
	setCmd.Flags().StringVar(&serParams.Value, "val", "", "Variable value")

	rootCmd.AddCommand(setCmd)
}

func runSetCmd(cmd *cobra.Command, args []string) error {
	if serParams.Environment == "" {
		return errors.New("invalid environment name")
	}

	if serParams.Variable == "" {
		return errors.New("invalid variable name")
	}

	variablesStore, err := store.New(configFile)
	if err != nil {
		return err
	}

	return variablesStore.Set(serParams.Environment, serParams.Variable, serParams.Value)
}
