package executor

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

// MockCommandRunner is a mock implementation of CommandRunner for testing.
type MockCommandRunner struct {
	// ExpectedCommand is the command name we expect to receive
	ExpectedCommand string
	// ExpectedArgs are the arguments we expect to receive
	ExpectedArgs []string
	// ReturnExitCode is the exit code to return
	ReturnExitCode int
	// ReturnError is the error to return (if any)
	ReturnError error
	// StdoutOutput is what to write to stdout
	StdoutOutput string
	// StderrOutput is what to write to stderr
	StderrOutput string
	// ActualCommand stores the command that was actually called
	ActualCommand string
	// ActualArgs stores the args that were actually passed
	ActualArgs []string
}

// Run implements the CommandRunner interface for testing.
func (m *MockCommandRunner) Run(name string, args []string, _ io.Reader, stdout, stderr io.Writer) (int, error) {
	m.ActualCommand = name
	m.ActualArgs = args

	if m.StdoutOutput != "" {
		fmt.Fprint(stdout, m.StdoutOutput)
	}
	if m.StderrOutput != "" {
		fmt.Fprint(stderr, m.StderrOutput)
	}

	return m.ReturnExitCode, m.ReturnError
}

// captureOutput is a helper that wraps ExecuteChezmoiWithRunner to capture stdout/stderr
// instead of printing to the console during tests.
func captureOutput(runner CommandRunner, args []string, useSudo bool) (exitCode int, stdout, stderr string, err error) {
	var outBuf, errBuf bytes.Buffer

	// Create a wrapper that captures output
	wrapper := &outputCapturingRunner{
		runner: runner,
		stdout: &outBuf,
		stderr: &errBuf,
	}

	exitCode, err = ExecuteChezmoiWithRunner(wrapper, args, useSudo)
	return exitCode, outBuf.String(), errBuf.String(), err
}

// outputCapturingRunner wraps a CommandRunner and redirects output to buffers
type outputCapturingRunner struct {
	runner CommandRunner
	stdout io.Writer
	stderr io.Writer
}

func (o *outputCapturingRunner) Run(name string, args []string, stdin io.Reader, _, _ io.Writer) (int, error) {
	return o.runner.Run(name, args, stdin, o.stdout, o.stderr)
}

// TestExecuteChezmoiWithRunner_BasicExecution tests successful command execution without sudo
func TestExecuteChezmoiWithRunner_BasicExecution(t *testing.T) {
	mock := &MockCommandRunner{
		ReturnExitCode: 0,
		StdoutOutput:   "chezmoi version v2.67.0\n",
	}

	exitCode, stdout, stderr, err := captureOutput(mock, []string{"--version"}, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	if mock.ActualCommand != "chezmoi" {
		t.Errorf("Expected command 'chezmoi', got '%s'", mock.ActualCommand)
	}

	if len(mock.ActualArgs) != 1 || mock.ActualArgs[0] != "--version" {
		t.Errorf("Expected args ['--version'], got %v", mock.ActualArgs)
	}

	if !strings.Contains(stdout, "chezmoi version") {
		t.Errorf("Expected stdout to contain version info, got: %s", stdout)
	}

	if stderr != "" {
		t.Errorf("Expected empty stderr, got: %s", stderr)
	}
}

// TestExecuteChezmoiWithRunner_WithSudo tests command execution with sudo flag
func TestExecuteChezmoiWithRunner_WithSudo(t *testing.T) {
	mock := &MockCommandRunner{
		ReturnExitCode: 0,
		StdoutOutput:   "chezmoi version v2.67.0\n",
	}

	exitCode, stdout, _, err := captureOutput(mock, []string{"--version"}, true)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	if mock.ActualCommand != "sudo" {
		t.Errorf("Expected command 'sudo', got '%s'", mock.ActualCommand)
	}

	// When using sudo, args should be ["chezmoi", "--version"]
	if len(mock.ActualArgs) != 2 || mock.ActualArgs[0] != "chezmoi" || mock.ActualArgs[1] != "--version" {
		t.Errorf("Expected args ['chezmoi', '--version'], got %v", mock.ActualArgs)
	}

	if !strings.Contains(stdout, "chezmoi version") {
		t.Errorf("Expected stdout to contain version info, got: %s", stdout)
	}
}

// TestExecuteChezmoiWithRunner_NonZeroExitCode tests handling of non-zero exit codes
func TestExecuteChezmoiWithRunner_NonZeroExitCode(t *testing.T) {
	mock := &MockCommandRunner{
		ReturnExitCode: 1,
		StderrOutput:   "chezmoi: unknown command\n",
	}

	exitCode, _, stderr, err := captureOutput(mock, []string{"invalid-command"}, false)
	if err != nil {
		t.Errorf("Expected no error from ExecuteChezmoi, got: %v", err)
	}

	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}

	if !strings.Contains(stderr, "unknown command") {
		t.Errorf("Expected stderr to contain 'unknown command', got: %s", stderr)
	}
}

// TestExecuteChezmoiWithRunner_CommandError tests handling of command execution errors
func TestExecuteChezmoiWithRunner_CommandError(t *testing.T) {
	mock := &MockCommandRunner{
		ReturnExitCode: -1,
		ReturnError:    fmt.Errorf("failed to execute command: exec: \"chezmoi\": executable file not found in $PATH"),
	}

	exitCode, _, _, err := captureOutput(mock, []string{"--version"}, false)

	if err == nil {
		t.Error("Expected error when command fails to execute, got nil")
	}

	if exitCode != -1 {
		t.Errorf("Expected exit code -1, got %d", exitCode)
	}
}

// TestExecuteChezmoiWithRunner_MultipleArgs tests passing multiple arguments
func TestExecuteChezmoiWithRunner_MultipleArgs(t *testing.T) {
	mock := &MockCommandRunner{
		ReturnExitCode: 0,
		StdoutOutput:   "add /etc/hosts\n",
	}

	args := []string{"add", "/etc/hosts"}
	exitCode, stdout, _, err := captureOutput(mock, args, false)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	if len(mock.ActualArgs) != 2 || mock.ActualArgs[0] != "add" || mock.ActualArgs[1] != "/etc/hosts" {
		t.Errorf("Expected args ['add', '/etc/hosts'], got %v", mock.ActualArgs)
	}

	if !strings.Contains(stdout, "/etc/hosts") {
		t.Errorf("Expected stdout to contain '/etc/hosts', got: %s", stdout)
	}
}

// TestExecuteChezmoiWithRunner_StdoutStderr tests that stdout and stderr are properly passed through
func TestExecuteChezmoiWithRunner_StdoutStderr(t *testing.T) {
	// We can't easily test that os.Stdout/os.Stderr are used in the main function,
	// but we can verify the mock receives writers and uses them
	var stdout, stderr bytes.Buffer

	// Create a custom runner to capture output
	customRunner := &MockCommandRunner{
		ReturnExitCode: 0,
		StdoutOutput:   "test output",
		StderrOutput:   "test error",
	}

	exitCode, err := customRunner.Run("chezmoi", []string{"--version"}, nil, &stdout, &stderr)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}

	if stdout.String() != "test output" {
		t.Errorf("Expected stdout 'test output', got '%s'", stdout.String())
	}

	if stderr.String() != "test error" {
		t.Errorf("Expected stderr 'test error', got '%s'", stderr.String())
	}
}
