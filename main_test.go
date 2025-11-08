package main

import (
	"io"
	"testing"
)

// mockRunnerMinimal is a lightweight mock for testing run().
type mockRunnerMinimal struct {
	code int
	err  error
}

func (m *mockRunnerMinimal) Run(name string, args []string, stdin io.Reader, stdout, stderr io.Writer) (int, error) {
	return m.code, m.err
}

// TestRun_Success covers the success path through run().
func TestRun_Success(t *testing.T) {
	defaultRunner = &mockRunnerMinimal{code: 0, err: nil}
	code := run([]string{"--version"})
	if code != 0 {
		// If chezmoi returned 0 we expect pass-through.
		t.Errorf("expected exit code 0, got %d", code)
	}
}

// TestRun_Error covers the error path mapping to exit code 1.
func TestRun_Error(t *testing.T) {
	defaultRunner = &mockRunnerMinimal{code: -1, err: errSentinel}
	code := run([]string{"--version"})
	if code != 1 {
		// Error should map to 1 irrespective of -1 sentinel.
		t.Errorf("expected exit code 1, got %d", code)
	}
}

// errSentinel is a static error used for testing.
var errSentinel = &sentinelError{}

// sentinelError implements error for test purposes.
type sentinelError struct{}

func (e *sentinelError) Error() string { return "sentinel" }
