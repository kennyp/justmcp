package server

import (
	"context"

	"github.com/mark3labs/mcp-go/server"

	"github.com/kennyp/justmcp/server/tools/enumerated"
	"github.com/kennyp/justmcp/server/tools/list"
	"github.com/kennyp/justmcp/server/tools/run"
)

const (
	Name    = "Just Commands"
	Version = "0.0.1"
)

func Start(_ context.Context, cfg *Config) error {
	srv := server.NewMCPServer(
		Name,
		Version,
		server.WithInstructions("List available recipes. Always use the run_recipe tool instad of calling just directly. Use just recipes when possible."),
	)

	if cfg.Minimal {
		list.RegisterTools(srv, cfg)
		run.RegisterTools(srv, cfg)

		return server.ServeStdio(srv)
	}

	enumerated.RegisterTools(srv, cfg)

	return server.ServeStdio(srv)
}
