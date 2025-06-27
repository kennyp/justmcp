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
	Exec(ctx context.Context, justfile string, args ...string) (*bytes.Buffer, error)
}

func RegisterTools(reg Registry, cfg Configurator) {
	listTool := mcp.NewTool(
		"list_recipes",
		mcp.WithDescription("List available recipes"),
		mcp.WithString("justfile",
			mcp.Description("Full path to the justfile for the project (ex. /home/kennyp/project/justfile)"),
			mcp.Required(),
		),
	)

	reg.AddTool(listTool, newHandler(cfg))
}

type listArgs struct {
	Justfile string `json:"justfile"`
}

type dump struct {
	Recipes map[string]*recipe `json:"recipes"`
}

type recipe struct {
	Doc string `json:"doc"`
}

func newHandler(cfg Configurator) server.ToolHandlerFunc {
	return mcp.NewTypedToolHandler(func(ctx context.Context, _ mcp.CallToolRequest, args listArgs) (*mcp.CallToolResult, error) {
		if args.Justfile == "" {
			return mcp.NewToolResultError("a justfile path is required"), nil
		}

		out, err := cfg.Exec(ctx, args.Justfile, "--list", "--color", "never")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("listing failed", err), nil
		}

		return mcp.NewToolResultText(out.String()), nil
	})
}
