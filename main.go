package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/kennyp/justmcp/server"
)

func main() {
	cmd := &cli.Command{
		Name:    "justmcp",
		Version: server.Version,
		Flags: []cli.Flag{
			&cli.BoolWithInverseFlag{
				Name:    "chdir",
				Usage:   "cd to the same file directory as justfile when running",
				Value:   true,
				Sources: cli.EnvVars("JUSTMCP_CHDIR"),
			},
			&cli.BoolWithInverseFlag{
				Name:    "mise",
				Aliases: []string{"m"},
				Usage:   "use 'mise x' when running just recipes",
				Sources: cli.EnvVars("JUSTMCP_MISE"),
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return server.Start(ctx, &server.Config{
				UseMise: c.Bool("mise"),
				Chdir:   c.Bool("chdir"),
			})
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("failed to start app", slog.Any("error", err))
		os.Exit(1)
	}
}
