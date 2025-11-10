package main

import (
	"os"

	"github.com/main-branch/chezroot/internal/executor"
)

// run is an extraction of main logic for testability. It returns the exit code
// that should be used by the process. Any error from ExecuteChezmoi is mapped
// to exit code 1 after writing its message to stderr.
func run(args []string) int {
	exitCode, err := executor.ExecuteChezmoi(args, false)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return 1
	}
	return exitCode
}

// exitFunc is a hook to allow tests to intercept process exit without terminating
// the test binary. It defaults to os.Exit.
var exitFunc = os.Exit

func main() {
	// Get command line arguments (excluding the program name itself)
	args := os.Args[1:]
	code := run(args)
	exitFunc(code)
}
