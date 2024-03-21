package cmd

import (
	"github.com/RevittConsulting/chain-dev-utils/internal/server"
	"github.com/spf13/cobra"
)

func Serve(server *server.Server) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "RunAsync Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			if err := server.Init(); err != nil {
				return err
			}
			return server.Run(ctx)
		},
	}
}
