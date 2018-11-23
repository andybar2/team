package main

import (
	"fmt"
	"os"

	"github.com/andybar2/team-env/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
