package list

import (
	"fmt"

	"github.com/spf13/cobra"

	items "github.com/crosleyzack/price_tracker/internal/items/fsjson"
)

// NewListItemsCommand creates a new command to add an item
func NewListItemsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "items",
		Example: "list items",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := items.NewConfig()
			if err != nil {
				return fmt.Errorf("failed to create item store config: %w", err)
			}
			itemStore, err := items.New(conf)
			if err != nil {
				return fmt.Errorf("failed to create item store: %w", err)
			}
			items, err := itemStore.ListItems()
			if err != nil {
				return fmt.Errorf("failed to list items: %w", err)
			}
			for _, item := range items {
				fmt.Println(item.String())
			}
			return nil
		},
	}
	return cmd
}
