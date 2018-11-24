package cmd

import (
	"errors"
	"fmt"

	"github.com/andybar2/team-env/store"
	"github.com/spf13/cobra"
)

var getParams struct {
	Environment string
	Variable    string
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an environment variable",
	RunE:  runGetCmd,
}

func init() {
	getCmd.Flags().StringVar(&getParams.Environment, "env", "", "Environment name")
	getCmd.Flags().StringVar(&getParams.Variable, "var", "", "Variable name")

	rootCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	if getParams.Environment == "" {
		return errors.New("invalid environment name")
	}

	if getParams.Variable == "" {
		return errors.New("invalid variable name")
	}

	variablesStore, err := store.New(configFile)
	if err != nil {
		return err
	}

	value, err := variablesStore.Get(getParams.Environment, getParams.Variable)
	if err != nil {
		return err
	}

	fmt.Printf("%s=%s\n", getParams.Variable, value)

	return nil
}
