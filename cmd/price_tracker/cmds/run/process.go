package run

import (
	"github.com/crosleyzack/price_tracker/internal/service"
	"github.com/spf13/cobra"
)

// NewProcessCommand creates a new command to launch the process
func NewProcessCommand() *cobra.Command {
	var serv *service.Service
	cmd := &cobra.Command{
		Use:     "process",
		Example: "run process",
		Args:    cobra.ExactArgs(0),
		PreRunE: func(_ *cobra.Command, _ []string) error {
			conf, err := service.NewConfig()
			if err != nil {
				return err
			}
			serv, err = service.NewService(conf)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return serv.NewEvents()
		},
	}
	return cmd
}
