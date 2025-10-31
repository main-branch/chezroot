//go:build tools

// This package imports internal development tools.
// It's not part of the main build.
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/rhysd/actionlint/cmd/actionlint"
)
