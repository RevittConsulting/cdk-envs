package cmd

import (
	"context"
	"fmt"
	"github.com/RevittConsulting/cdk-envs/config"
	"github.com/RevittConsulting/cdk-envs/internal/server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func RootCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:              "root",
		Short:            "root is the root command for the application",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfgLoad, err := config.LoadConfig()
			if err != nil {
				return err
			}
			*cfg = *cfgLoad

			return nil
		},
	}
}

func Root() {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	cfg := &config.Config{}
	rootCmd := RootCmd(cfg)

	svr := &server.Server{
		Config: cfg,
	}

	rootCmd.AddCommand(Serve(svr))

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
