package list

import (
	"fmt"

	"github.com/spf13/cobra"

	events "github.com/crosleyzack/price_tracker/internal/events/fsjson"
	"github.com/crosleyzack/price_tracker/internal/model"
)

// NewListEventsCommand creates a new command to add an item
func NewListEventsCommand() *cobra.Command {
	var itemName string
	cmd := &cobra.Command{
		Use:     "events",
		Example: "list events",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := events.NewConfig()
			if err != nil {
				return fmt.Errorf("failed to create item store config: %w", err)
			}
			eventStore, err := events.New(conf)
			if err != nil {
				return fmt.Errorf("failed to create item store: %w", err)
			}
			// if itemName is not empty, list events for that item
			// otherwise, list all events
			var events []model.Event
			if itemName != "" {
				events, err = eventStore.ListEventsForItem(itemName)
			} else {
				events, err = eventStore.ListEvents()
			}
			if err != nil {
				return fmt.Errorf("failed to list events: %w", err)
			}
			for _, event := range events {
				fmt.Println(event.String())
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&itemName, "item", "i", "", "item to list events for")
	return cmd
}
