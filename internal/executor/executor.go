package executor

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// CommandRunner is an interface for running commands.
// This allows us to mock command execution in tests.
type CommandRunner interface {
	Run(name string, args []string, stdin io.Reader, stdout, stderr io.Writer) (int, error)
}

// RealCommandRunner implements CommandRunner using os/exec.
type RealCommandRunner struct{}

// Run executes a command with the given arguments and returns the exit code.
func (r *RealCommandRunner) Run(name string, args []string, stdin io.Reader, stdout, stderr io.Writer) (int, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()

	exitCode := 0
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			exitCode = exitErr.ExitCode()
		} else {
			// If we couldn't run the command at all (e.g., command not found)
			return -1, fmt.Errorf("failed to execute command: %w", err)
		}
	}

	return exitCode, nil
}

// defaultRunner is the global default command runner.
var defaultRunner CommandRunner = &RealCommandRunner{}

// GetDefaultRunner returns the current default runner (for testing).
func GetDefaultRunner() CommandRunner {
	return defaultRunner
}

// SetDefaultRunner sets the default runner (for testing).
func SetDefaultRunner(runner CommandRunner) {
	defaultRunner = runner
}

// ExecuteChezmoi runs chezmoi in a subprocess with the given arguments.
// If useSudo is true, it prefixes the command with "sudo".
// stdin, stdout, and stderr are passed through transparently.
func ExecuteChezmoi(args []string, useSudo bool) (int, error) {
	return ExecuteChezmoiWithRunner(defaultRunner, args, useSudo)
}

// ExecuteChezmoiWithRunner runs chezmoi using the provided CommandRunner.
// This function is useful for testing with a mock runner.
func ExecuteChezmoiWithRunner(runner CommandRunner, args []string, useSudo bool) (int, error) {
	var cmdName string
	var cmdArgs []string

	if useSudo {
		cmdName = "sudo"
		// Prepend "chezmoi" to the arguments
		cmdArgs = append([]string{"chezmoi"}, args...)
	} else {
		cmdName = "chezmoi"
		cmdArgs = args
	}

	return runner.Run(cmdName, cmdArgs, os.Stdin, os.Stdout, os.Stderr)
}
