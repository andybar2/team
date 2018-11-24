package cmd

import (
	"errors"

	"github.com/andybar2/team-env/store"
	"github.com/spf13/cobra"
)

var delParams struct {
	Environment string
	Variable    string
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an environment variable",
	RunE:  runDelCmd,
}

func init() {
	delCmd.Flags().StringVar(&delParams.Environment, "env", "", "Environment name")
	delCmd.Flags().StringVar(&delParams.Variable, "var", "", "Variable name")

	rootCmd.AddCommand(delCmd)
}

func runDelCmd(cmd *cobra.Command, args []string) error {
	if delParams.Environment == "" {
		return errors.New("invalid environment name")
	}

	if delParams.Variable == "" {
		return errors.New("invalid variable name")
	}

	variablesStore, err := store.New(configFile)
	if err != nil {
		return err
	}

	return variablesStore.Del(delParams.Environment, delParams.Variable)
}
