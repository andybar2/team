package env

import (
	"errors"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var printParams struct {
	Stage string
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print all the environment variables for a stage",
	RunE:  runPrintCmd,
}

func init() {
	printCmd.Flags().StringVarP(&printParams.Stage, "stage", "s", "", "Stage name")

	EnvCmd.AddCommand(printCmd)
}

func runPrintCmd(cmd *cobra.Command, args []string) error {
	if printParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	return s.EnvPrint(printParams.Stage)
}
