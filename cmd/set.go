package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var setEnvironment, setVariable, setValue string

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an environment variable",
	RunE:  runSetCmd,
}

func init() {
	setCmd.Flags().StringVar(&setEnvironment, "env", "", "Environment name")
	setCmd.Flags().StringVar(&setVariable, "var", "", "Variable name")
	setCmd.Flags().StringVar(&setValue, "val", "", "Variable value")

	rootCmd.AddCommand(setCmd)
}

func runSetCmd(cmd *cobra.Command, args []string) error {
	if setEnvironment == "" {
		return errors.New("invalid environment name")
	}

	if setVariable == "" {
		return errors.New("invalid variable name")
	}

	variablesStore, err := setupStore()
	if err != nil {
		return err
	}

	return variablesStore.Set(setEnvironment, setVariable, setValue)
}
