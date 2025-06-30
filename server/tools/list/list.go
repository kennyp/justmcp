package list

import (
	"bytes"
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Registry interface {
	AddTool(mcp.Tool, server.ToolHandlerFunc)
}

type Configurator interface {
	Exec(ctx context.Context, args ...string) (*bytes.Buffer, error)
}

func RegisterTools(reg Registry, cfg Configurator) {
	listTool := mcp.NewTool(
		"list_recipes",
		mcp.WithDescription("List available recipes"),
	)

	reg.AddTool(listTool, newHandler(cfg))
}

type listArgs struct {
}

func newHandler(cfg Configurator) server.ToolHandlerFunc {
	return mcp.NewTypedToolHandler(func(ctx context.Context, _ mcp.CallToolRequest, args listArgs) (*mcp.CallToolResult, error) {
		out, err := cfg.Exec(ctx, "--list", "--color", "never")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("listing failed", err), nil
		}

		return mcp.NewToolResultText(out.String()), nil
	})
}
