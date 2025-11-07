# Contributing to chezroot

First, thank you for considering contributing to `chezroot`. This project is a
community effort, and every contribution is valued.

This guide outlines the standards and procedures for contributing, from setting up
your environment to submitting your changes. Please read it carefully to ensure your
contributions can be merged smoothly.

- [TLDR](#tldr)
- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Environment Setup](#development-environment-setup)
- [Development Workflow](#development-workflow)
- [Submitting Changes](#submitting-changes)
- [Automation Pipeline](#automation-pipeline)

## TLDR

Here are the most important guidelines:

- **Prerequisites:** You must have Go (>= 1.25.1), Node.js (>= 24), and npm installed.
- **One-Time Setup:** Run `make install-tools` after cloning to install all
  development dependencies and Git hooks.
- **Conventional Commits:** All commit messages *must* follow the Conventional
  Commits standard (e.g., `feat: ...`, `fix: ...`). A local Git hook will enforce
  this.
- **Tests are Required:** All new features or bug fixes must include corresponding
  tests.
- **Update Documentation:** Any change that impacts users (new features, flags, etc.)
  **must** be documented appropriately
- **Ensure CI passes locally:** Before submitting a pull request, run `make ci` and
  ensure all checks (linting, testing, building) pass. Note: this will fail if
  `go mod tidy` makes changes to `go.mod` or `go.sum`; commit those changes.
- **Rebase Workflow:** Always work on a new branch. Your pull request **must be
  rebased** on the latest `main` branch before it will be merged with a
  fast-forward merge.

## Code of Conduct

This project and everyone participating in it is governed by the [Code of
Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

Before you begin, you'll need to fork and clone the repository.

1. **Fork** the repository on GitHub.

2. **Clone** your fork to your local machine:

   ```bash
   git clone https://github.com/YOUR_USERNAME/chezroot.git
   cd chezroot
   ```

3. **Add the upstream remote** to keep your fork in sync:

   ```bash
   git remote add upstream https://github.com/main-branch/chezroot.git
   ```

## Development Environment Setup

This project uses Go for the application and Node.js for development tooling
(linters, commit hooks).

### Prerequisites

You must have the following tools installed on your local machine:

- **Go:** Version 1.25.1 or newer.
- **Node.js:** Version 24.0.0 or newer.
- **npm:** This is included with Node.js.

### Install Dependencies and Tooling

Once you have the prerequisites, you can install all project dependencies and
development tools by running one command:

```bash
make install-tools
```

This command will:

1. Run `npm ci`, which installs all Node.js-based development tools (linters, commit
  hooks) from the [package-lock.json](package-lock.json) file.

2. Run the `prepare` script in [package.json](package.json), which uses Husky to set up the
  `commit-msg` Git hook ([.husky/commit-msg](.husky/commit-msg)).

The Go-based development tools (`golangci-lint`, `actionlint`, `goreleaser`) are
managed via [go.mod](go.mod) and invoked with `go run` via the
[Makefile](Makefile). You do not need to install them globally.

If you later pull changes that modify
[package-lock.json](package-lock.json), re-run `make install-tools` to keep local
development tooling in sync.

### Editor Setup (VS Code)

For the best development experience, we recommend using VS Code with the following
extensions:

- **Go:** `golang.Go`
- **ESLint:** `dbaeumer.vscode-eslint`
- **Markdownlint:** `DavidAnson.vscode-markdownlint`

This repository includes a [.vscode/settings.json](.vscode/settings.json) file that will automatically
configure these extensions to use the project's rules (e.g., formatting your YAML
files on save to match our [eslint.config.js](eslint.config.js) rules).

## Development Workflow

The `Makefile` is the single source of truth for all common development tasks.

### Main Command

Before submitting a pull request, **you must run the main CI check** locally:

```bash
make ci
```

This command runs the full suite of linters, tests, builds the binary, and checks
that your Go modules are tidy. This is the exact same command our GitHub workflow
uses.

Note: the CI check will fail if `go mod tidy` results in changes to
[go.mod](go.mod) or [go.sum](go.sum). If that happens locally, please commit those
changes in your branch before
opening or updating a pull request.

### Other Useful Commands

Run `make help` to see all available development targets (from the
[Makefile](Makefile)) and their descriptions.

## Submitting Changes

### Commit Message Format

This project uses **Conventional Commits**. This format is strictly enforced by a Git
hook and our CI pipeline.

For the full specification and examples, see the
[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) website.

After running `make install-tools`, a `commit-msg` hook is installed that
automatically lints your commit message with `commitlint`.

- Your commit message **must** follow the format: `<type>(<scope>): <subject>`.
- Allowed types are: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`,
  `build`, `ci`, `chore`, `revert`.
- For details on all rules, see [./.commitlintrc.yml](./.commitlintrc.yml).

**Examples:**

- `feat(profile): add new 'purge' command`
- `fix: correct sudo handling for edit command`
- `docs: update installation instructions in README`

### Pull Request Process

1. Create a feature branch from the `main` branch.
2. Make your changes.

      **Important Contributor Requirements**

      - **All new code requires tests:** Any new feature or bug fix must be accompanied by
        corresponding tests.

      - **Documentation must be updated:** Any change that impacts user behavior (new
        commands, flags, or profile changes) must be documented in the
        [README.md](README.md).

3. Ensure your code passes all local checks by running `make ci`.
4. Rebase your branch on the latest `main`:

    ```bash
    git fetch upstream
    git rebase upstream/main
    ```

    If you encounter conflicts during the rebase, resolve them and continue:

    ```bash
    git add <files>
    git rebase --continue
    ```

5. Push your branch and open a pull request against `main-branch/chezroot:main`.
6. Ensure all CI checks on the pull request pass.

## Automation Pipeline

We use a set of GitHub Actions to automate linting, testing, and releases.

- **Continuous Integration:** On every pull request,
  [.github/workflows/continuous-integration.yml](.github/workflows/continuous-integration.yml)
  runs `make ci` to ensure all tests and linters pass.
- **Conventional Commits:** The
  [.github/workflows/enforce-conventional-commits.yml](.github/workflows/enforce-conventional-commits.yml)
  workflow lints all commit messages in the pull request.
- **Dependency Updates:**
  [.github/dependabot.yml](.github/dependabot.yml) is configured to open pull
  requests for Go and npm dependency updates.
- **Release Automation:**
  1. When a pull request is merged to `main`,
     [.github/workflows/create-release.yml](.github/workflows/create-release.yml)
     runs `release-please` to open (or update) a "Release PR" with the new version
     and [CHANGELOG.md](CHANGELOG.md).
  2. When that "Release PR" is merged, a new version tag is created.
  3. This tag triggers
     [.github/workflows/publish-release.yml](.github/workflows/publish-release.yml),
     which uses `goreleaser` to build binaries, upload them, and update the
     Homebrew tap.
