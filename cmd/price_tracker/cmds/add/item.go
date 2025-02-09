package add

import (
	"fmt"

	"github.com/spf13/cobra"

	items "github.com/crosleyzack/price_tracker/internal/items/fsjson"
	"github.com/crosleyzack/price_tracker/internal/model"
)

// NewAddItemCommand creates a new command to add an item
func NewAddItemCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "item",
		Example: "add item <name> <website> <uri-path>",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := items.NewConfig()
			if err != nil {
				return fmt.Errorf("failed to create item store config: %w", err)
			}
			itemStore, err := items.New(conf)
			if err != nil {
				return fmt.Errorf("failed to create item store: %w", err)
			}
			item := model.Item{
				Name:    args[0],
				Website: args[1],
				URIPath: args[2],
			}
			err = itemStore.AddItem(item)
			if err != nil {
				return fmt.Errorf("failed to add item: %w", err)
			}
			fmt.Printf("Item added: %v\n", item)
			return nil
		},
	}
	return cmd
}
