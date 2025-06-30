package enumerated

import (
	"bytes"
	"context"
	"fmt"
	"iter"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/kennyp/justmcp/parser"
)

type Registry interface {
	AddTool(mcp.Tool, server.ToolHandlerFunc)
}

type Configurator interface {
	RecipesToRegister() iter.Seq2[string, *parser.Recipe]
	Exec(ctx context.Context, args ...string) (*bytes.Buffer, error)
}

func RegisterTools(reg Registry, cfg Configurator) {
	for name, recipe := range cfg.RecipesToRegister() {
		opts := []mcp.ToolOption{
			mcp.WithDescription(recipe.Doc.Description()),
		}

		parsers := make([]ParamParser, len(recipe.Parameters))
		for i, p := range recipe.Parameters {
			switch p.Kind {
			case "singular":
				parsers[i] = SingularParser(p.Name)
				opts = append(opts, mcp.WithString(
					p.Name,
					mcp.Description(p.Doc),
				))
			case "plus":
				parsers[i] = PlusParser(p.Name)
				opts = append(opts, mcp.WithArray(
					p.Name,
					mcp.Items(map[string]any{"type": "string"}),
					mcp.Description(p.Doc),
				))
			}
		}

		reg.AddTool(mcp.NewTool(name, opts...), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := []string{name}
			for _, p := range parsers {
				pargs, err := p.Parse(ctx, request)
				if err != nil {
					return mcp.NewToolResultErrorFromErr("bad argument", err), nil
				}

				args = append(args, pargs...)
			}

			out, err := cfg.Exec(ctx, args...)
			if err != nil {
				return mcp.NewToolResultErrorFromErr(fmt.Sprintf("%s failed", name), err), nil
			}

			return mcp.NewToolResultText(out.String()), nil
		})
	}
}

type ParamParser interface {
	Parse(context.Context, mcp.CallToolRequest) ([]string, error)
}

type SingularParser string

func (p SingularParser) Parse(_ context.Context, request mcp.CallToolRequest) ([]string, error) {
	return []string{request.GetString(string(p), "")}, nil
}

type PlusParser string

func (p PlusParser) Parse(_ context.Context, request mcp.CallToolRequest) ([]string, error) {
	return request.GetStringSlice(string(p), []string{}), nil
}
