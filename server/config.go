package server

import (
	"bytes"
	"context"
	"fmt"
	"iter"
	"os"
	"os/exec"
	"path"
	"slices"

	"github.com/kennyp/justmcp/parser"
)

type Config struct {
	Justfile     *parser.Justfile
	UseMise      bool
	Chdir        bool
	Minimal      bool
	AllowedTools []string
}

func (c *Config) Allowed(name string) bool {
	if slices.Index(c.AllowedTools, name) == -1 {
		return len(c.AllowedTools) == 0
	}

	return true
}

func (c *Config) RecipesToRegister() iter.Seq2[string, *parser.Recipe] {
	return func(yield func(string, *parser.Recipe) bool) {
		for name, recipe := range c.Justfile.PublicRecipes() {
			if !c.Allowed(name) {
				continue
			}

			if !yield(name, recipe) {
				return
			}
		}
	}
}

func (c *Config) Command() string {
	if c.UseMise {
		return "mise"
	}

	return "just"
}

func (c *Config) BaseArgs() []string {
	if c.UseMise {
		return []string{"x", "--", "just", "-f", c.Justfile.Path}
	}

	return []string{"-f", c.Justfile.Path}
}

func (c *Config) Exec(ctx context.Context, args ...string) (*bytes.Buffer, error) {
	if c.Chdir {
		if err := os.Chdir(path.Dir(c.Justfile.Path)); err != nil {
			return nil, fmt.Errorf("failed to change dir (%w)", err)
		}
	}

	cmdArgs := append(c.BaseArgs(), args...)
	cmd := exec.CommandContext(ctx, c.Command(), cmdArgs...)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	return &out, cmd.Run()
}
