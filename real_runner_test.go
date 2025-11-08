package main

import (
	"bytes"
	"testing"
)

// TestRealCommandRunner_Success verifies successful execution with output.
func TestRealCommandRunner_Success(t *testing.T) {
	runner := &RealCommandRunner{}
	var stdout, stderr bytes.Buffer

	exitCode, err := runner.Run("bash", []string{"-c", "echo hi"}, nil, &stdout, &stderr)
	if err != nil {
		// RealCommandRunner suppresses exitError when exit code is 0; err should be nil
		// Any non-nil error here is unexpected.
		// Provide detail to help debugging environment issues.
		// Use t.Fatalf so subsequent assertions don't run.
		// (E.g. missing bash on minimal systems.)
		// On macOS bash should exist.
		//
		// We intentionally do not skip this test; it's fundamental.
		// If needed, adjust command to /bin/sh.
		//
		t.Fatalf("expected nil error, got %v", err)
	}
	if exitCode != 0 {
		t.Errorf("expected exit code 0, got %d", exitCode)
	}
	if stdout.String() != "hi\n" {
		t.Errorf("expected stdout 'hi\\n', got %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Errorf("expected empty stderr, got %q", stderr.String())
	}
}

// TestRealCommandRunner_NonZeroExit verifies non-zero exit captures exit code and yields nil error.
func TestRealCommandRunner_NonZeroExit(t *testing.T) {
	runner := &RealCommandRunner{}
	var stdout, stderr bytes.Buffer
	exitCode, err := runner.Run("bash", []string{"-c", "exit 7"}, nil, &stdout, &stderr)
	if err != nil {
		t.Fatalf("expected nil error for non-zero exit, got %v", err)
	}
	if exitCode != 7 {
		t.Errorf("expected exit code 7, got %d", exitCode)
	}
}

// TestRealCommandRunner_CommandNotFound verifies handling of an execution error (command not found).
func TestRealCommandRunner_CommandNotFound(t *testing.T) {
	runner := &RealCommandRunner{}
	var stdout, stderr bytes.Buffer
	exitCode, err := runner.Run("___no_such_command___", nil, nil, &stdout, &stderr)
	if err == nil {
		t.Fatalf("expected error for unknown command, got nil")
	}
	if exitCode != -1 {
		t.Errorf("expected exit code -1, got %d", exitCode)
	}
}
