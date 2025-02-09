package run

import (
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra command to run a target
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "",
	}
	cmd.AddCommand(NewProcessCommand())
	return cmd
}
