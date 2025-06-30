package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/yassinebenaid/godump"

	"github.com/kennyp/justmcp/parser"
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
			&cli.StringFlag{
				Name:     "justfile",
				Aliases:  []string{"f"},
				Usage:    "path to `justfile`",
				Sources:  cli.EnvVars("JUSTMCP_JUSTFILE", "JUST_JUSTFILE"),
				Required: true,
				Validator: func(p string) (err error) {
					_, err = os.Stat(p)
					return
				},
			},
			&cli.BoolWithInverseFlag{
				Name:    "minimal",
				Usage:   "only register minimal tools",
				Sources: cli.EnvVars("JUSTMCP_MINIMAL"),
			},
			&cli.StringSliceFlag{
				Name:    "tools",
				Aliases: []string{"t"},
				Usage:   "only allow the given `tools`",
				Sources: cli.EnvVars("JUSTMCP_TOOLS"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "dump",
				Usage: "dump justfile for debugging",
				Action: func(ctx context.Context, c *cli.Command) error {
					f, err := parser.Parse(ctx, c.String("justfile"))
					if err != nil {
						return err
					}

					return godump.Dump(f)
				},
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			f, err := parser.Parse(ctx, c.String("justfile"))
			if err != nil {
				return err
			}

			return server.Start(ctx, &server.Config{
				Justfile:     f,
				UseMise:      c.Bool("mise"),
				Chdir:        c.Bool("chdir"),
				Minimal:      c.Bool("minimal"),
				AllowedTools: c.StringSlice("tools"),
			})
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("failed to start app", slog.Any("error", err))
		os.Exit(1)
	}
}
