package cmds

import (
	"github.com/spf13/cobra"

	"github.com/crosleyzack/price_tracker/cmd/price_tracker/cmds/add"
	"github.com/crosleyzack/price_tracker/cmd/price_tracker/cmds/list"
	"github.com/crosleyzack/price_tracker/cmd/price_tracker/cmds/run"
)

// NewCommand creates a new cobra command
func NewCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(run.NewCommand(), add.NewCommand(), list.NewCommand())
	return rootCmd, nil
}
