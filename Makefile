# Makefile for the chezroot project
#
# This Makefile uses project-local dependencies defined in
# go.mod (via tools.go) and package.json.

.DEFAULT_GOAL := lint

# Pin linter versions to those in go.mod
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint
ACTIONLINT := go run github.com/rhysd/actionlint/cmd/actionlint
GORELEASER := go run github.com/goreleaser/goreleaser/v2

# Define all task targets as .PHONY
.PHONY: ci install-tools lint lint-fix lint-go lint-md lint-yaml lint-actions build test clean clean-all

# ==============================================================================
# CI (Continuous Integration)
# ==============================================================================

# ci: Run all checks required for CI
ci: lint test build
	@echo "--> Checking Go module tidy..."
	@go mod tidy -v
	@if ! git diff --exit-code -- go.mod go.sum; then \
		echo "FAILURE: 'go mod tidy' resulted in changes. Please commit them."; \
		exit 1; \
	fi

# ==============================================================================
# Setup
# ==============================================================================

# install-tools: Install local npm dependencies
install-tools:
	@echo "--> Installing dependencies from lockfiles..."
	@npm ci

# ==============================================================================
# Linting
# ==============================================================================

# lint: Run all linters
lint: lint-go lint-md lint-yaml lint-actions lint-goreleaser
	@echo "✅ All linters passed."

# lint-fix: Automatically fix all fixable linting errors
lint-fix:
	@echo "--> Fixing all lintable files..."
	@$(GOLANGCI_LINT) run --fix ./...
	@npx eslint . --fix

# lint-go: Run the Go linter
lint-go:
	@echo "--> Linting Go code..."
	@$(GOLANGCI_LINT) run ./...

# lint-md: Run the Markdown linter
lint-md:
	@echo "--> Linting Markdown files..."
	@npx markdownlint-cli2 "**/*.md" "#node_modules" "#CHANGELOG.md" > /dev/null

# lint-yaml: Run the YAML linter (using ESLint)
lint-yaml:
	@echo "--> Linting YAML files..."
	@npx eslint .

# lint-actions: Run the GitHub Actions linter
lint-actions:
	@echo "--> Linting GitHub Actions..."
	@$(ACTIONLINT) -color

# lint-goreleaser: Run the GoReleaser config check
lint-goreleaser:
	@echo "--> Linting GoReleaser config..."
	@$(GORELEASER) check --quiet

# ==============================================================================
# Build & Test
# ==============================================================================

# build: Compile the Go application
build:
	@echo "--> Building Go binary..."
	@go build -v ./...

# test: Run all Go tests
test:
	@echo "--> Running Go tests..."
	@go test -v -race ./...

# ==============================================================================
# Cleaning
# ==============================================================================

# clean: Remove compiled binaries and test/coverage files
clean:
	@echo "--> Cleaning build artifacts..."
	@rm -f chezroot
	@rm -rf ./bin
	@rm -f coverage.*
	@rm -f *.out
	@rm -rf ./dist

# clean-all: Remove all generated files, including local dependencies
clean-all: clean
	@echo "--> Clobbering all dependencies..."
	@rm -rf ./node_modules
	@rm -rf ./.husky/_