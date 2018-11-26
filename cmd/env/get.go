package env

import (
	"errors"
	"fmt"

	"github.com/andybar2/team/store"
	"github.com/spf13/cobra"
)

var getParams struct {
	Stage string
	Name  string
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an environment variable",
	RunE:  runGetCmd,
}

func init() {
	getCmd.Flags().StringVarP(&getParams.Stage, "stage", "s", "", "Stage name")
	getCmd.Flags().StringVarP(&getParams.Name, "name", "n", "", "Variable name")

	EnvCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	if getParams.Stage == "" {
		return errors.New("invalid stage name")
	}

	if getParams.Name == "" {
		return errors.New("invalid variable name")
	}

	s, err := store.New()
	if err != nil {
		return err
	}

	value, err := s.EnvGet(getParams.Stage, getParams.Name)
	if err != nil {
		return err
	}

	fmt.Printf("%s=%s\n", getParams.Name, value)

	return nil
}
