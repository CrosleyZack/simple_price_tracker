package add

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/crosleyzack/price_tracker/internal/model"
	sites "github.com/crosleyzack/price_tracker/internal/sites/fsjson"
)

// NewAddSiteCommand creates a new command to add an item
func NewAddSiteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "site",
		Example: "add site <name> <url> <price path>",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := sites.NewConfig()
			if err != nil {
				return fmt.Errorf("failed to create item store config: %w", err)
			}
			siteStore, err := sites.New(conf)
			if err != nil {
				return fmt.Errorf("failed to create item store: %w", err)
			}
			site := model.Website{
				Name:      args[0],
				URL:       args[1],
				PricePath: args[2],
			}
			err = siteStore.AddSite(site)
			if err != nil {
				return fmt.Errorf("failed to add site: %w", err)
			}
			fmt.Printf("Site added: %v\n", site)
			return nil
		},
	}
	return cmd
}
