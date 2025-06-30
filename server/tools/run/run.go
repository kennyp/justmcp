package run

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
	runTool := mcp.NewTool(
		"run_recipe",
		mcp.WithDescription("Execute a recipe"),
		mcp.WithString("recipe",
			mcp.Description("Name of the recipe to execute (ex. build)"),
			mcp.Required(),
		),
		mcp.WithArray("arguments",
			mcp.Description("additional arguments to the recipe"),
			mcp.Items(map[string]any{"type": "string"}),
		),
	)

	reg.AddTool(runTool, newHandler(cfg))
}

type executeArgs struct {
	Recipe string   `json:"recipe"`
	Args   []string `json:"arguments"`
}

func newHandler(cfg Configurator) server.ToolHandlerFunc {
	return mcp.NewTypedToolHandler(func(ctx context.Context, _ mcp.CallToolRequest, args executeArgs) (*mcp.CallToolResult, error) {
		if args.Recipe == "" {
			return mcp.NewToolResultError("a recipe name is required"), nil
		}

		allArgs := append([]string{args.Recipe}, args.Args...)

		out, err := cfg.Exec(ctx, allArgs...)
		if err != nil {
			return mcp.NewToolResultErrorFromErr(out.String(), err), nil
		}

		return mcp.NewToolResultText(out.String()), nil
	})
}
