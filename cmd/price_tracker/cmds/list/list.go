package list

import (
	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra command to run a target
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "",
	}
	cmd.AddCommand(NewListItemsCommand())
	cmd.AddCommand(NewListSitesCommand())
	cmd.AddCommand(NewListEventsCommand())
	return cmd
}
