package add

import (
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra command to run a target
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "",
	}
	cmd.AddCommand(NewAddItemCommand())
	cmd.AddCommand(NewAddSiteCommand())
	return cmd
}
