package main

import (
	"fmt"
	"os"

	"github.com/andybar2/team/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
