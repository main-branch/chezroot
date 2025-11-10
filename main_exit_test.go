package main

import (
	"io"
	"os"
	"testing"

	"github.com/main-branch/chezroot/internal/executor"
)

// mockExitRunner implements CommandRunner to drive main() paths.
type mockExitRunner struct {
	code int
	err  error
}

func (m *mockExitRunner) Run(_ string, _ []string, _ io.Reader, _, _ io.Writer) (int, error) {
	return m.code, m.err
}

// captureStderr replaces os.Stderr with a pipe and returns a function to restore it and the captured output.
func captureStderr(t *testing.T) (restore func() string) {
	orig := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stderr = w
	return func() string {
		w.Close()
		os.Stderr = orig
		bytes, _ := io.ReadAll(r)
		return string(bytes)
	}
}

// TestMainFunction_Success ensures main invokes exitFunc with underlying exit code.
func TestMainFunction_Success(t *testing.T) {
	// Save original runner
	originalRunner := executor.GetDefaultRunner()
	defer executor.SetDefaultRunner(originalRunner)

	executor.SetDefaultRunner(&mockExitRunner{code: 0, err: nil})
	var got int
	exitFunc = func(c int) { got = c }
	defer func() { exitFunc = os.Exit }()
	main()
	if got != 0 {
		t.Errorf("expected exit code 0, got %d", got)
	}
}

// TestMainFunction_Error ensures error path maps to exit code 1 and writes stderr.
func TestMainFunction_Error(t *testing.T) {
	// Save original runner
	originalRunner := executor.GetDefaultRunner()
	defer executor.SetDefaultRunner(originalRunner)

	executor.SetDefaultRunner(&mockExitRunner{code: -1, err: errSentinel})
	var got int
	exitFunc = func(c int) { got = c }
	defer func() { exitFunc = os.Exit }()
	restore := captureStderr(t)
	main()
	stderr := restore()
	if got != 1 {
		t.Errorf("expected exit code 1, got %d", got)
	}
	if stderr == "" {
		t.Errorf("expected stderr output, got empty string")
	}
	if stderr != "sentinel\n" { // run() adds newline
		// Provide len for easier debugging
		t.Errorf("unexpected stderr contents: %q (len=%d)", stderr, len(stderr))
	}
}
