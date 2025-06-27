package server

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Config struct {
	UseMise bool
	Chdir   bool
}

func (c *Config) Command() string {
	if c.UseMise {
		return "mise"
	}

	return "just"
}

func (c *Config) BaseArgs() []string {
	if c.UseMise {
		return []string{"x", "--", "just"}
	}

	return []string{}
}

func (c *Config) Exec(ctx context.Context, justfile string, args ...string) (*bytes.Buffer, error) {
	if c.Chdir {
		if err := os.Chdir(path.Dir(justfile)); err != nil {
			return nil, fmt.Errorf("failed to change dir (%w)", err)
		}
	}

	cmdArgs := append(c.BaseArgs(), "-f", justfile)
	cmdArgs = append(cmdArgs, args...)

	var out bytes.Buffer

	cmd := exec.CommandContext(ctx, c.Command(), cmdArgs...)
	cmd.Stdout = &out
	cmd.Stderr = &out

	return &out, cmd.Run()
}
