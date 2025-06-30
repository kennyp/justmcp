# Just Settings

set unstable := true
set script-interpreter := ['bash', '-eu', '-o', 'pipefail']

# Helper Commands

golangci-lint := "go tool golangci-lint"
modernize := "go tool modernize"

# Show this
@help:
    just --list --list-heading "" --justfile {{ justfile() }}

# Build all binaries
@build: lint
    go build

# Update vendor directory to match go.mod
@vendor:
    go mod tidy
    go mod vendor

# Lint check the code `package:"pacages to lint"`
@lint +package="./...":
    {{ golangci-lint }} run {{ package }}
    {{ modernize }} -test {{ package }}
