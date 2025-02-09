package list

import (
	"fmt"

	"github.com/spf13/cobra"

	sites "github.com/crosleyzack/price_tracker/internal/sites/fsjson"
)

// NewListSitesCommand creates a new command to add an item
func NewListSitesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sites",
		Example: "list sites",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := sites.NewConfig()
			if err != nil {
				return fmt.Errorf("failed to create item store config: %w", err)
			}
			siteStore, err := sites.New(conf)
			if err != nil {
				return fmt.Errorf("failed to create item store: %w", err)
			}
			sites, err := siteStore.ListSites()
			if err != nil {
				return fmt.Errorf("failed to list sites: %w", err)
			}
			for _, site := range sites {
				fmt.Println(site.String())
			}
			return nil
		},
	}
	return cmd
}
