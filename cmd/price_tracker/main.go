package main

import (
	"fmt"
	"os"

	"github.com/crosleyzack/price_tracker/cmd/price_tracker/cmds"
)

// main loads root command from cmds package and executes it
func main() {
	rootCmd, err := cmds.NewCommand()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
