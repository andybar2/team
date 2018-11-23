package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getEnvironment, getVariable string

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an environment variable",
	RunE:  runGetCmd,
}

func init() {
	getCmd.Flags().StringVar(&getEnvironment, "env", "", "Environment name")
	getCmd.Flags().StringVar(&getVariable, "var", "", "Variable name")

	rootCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	if getEnvironment == "" {
		return errors.New("invalid environment name")
	}

	if getVariable == "" {
		return errors.New("invalid variable name")
	}

	variablesStore, err := setupStore()
	if err != nil {
		return err
	}

	value, err := variablesStore.Get(getEnvironment, getVariable)
	if err != nil {
		return err
	}

	fmt.Printf("%s=%s\n", getVariable, value)

	return nil
}
