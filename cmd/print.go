package cmd

import (
	"errors"

	"github.com/andybar2/team-env/store"
	"github.com/spf13/cobra"
)

var printParams struct {
	Environment string
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all variables for an environment",
	RunE:  runPrintCmd,
}

func init() {
	printCmd.Flags().StringVar(&printParams.Environment, "env", "", "Environment name")

	rootCmd.AddCommand(printCmd)
}

func runPrintCmd(cmd *cobra.Command, args []string) error {
	if printParams.Environment == "" {
		return errors.New("invalid environment name")
	}

	variablesStore, err := store.New(configFile)
	if err != nil {
		return err
	}

	return variablesStore.Print(getParams.Environment)
}
