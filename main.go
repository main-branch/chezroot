package main

import (
	"os"
)

func main() {
	// Get command line arguments (excluding the program name itself)
	args := os.Args[1:]

	// For now, always pass useSudo as false as per issue #42
	exitCode, err := ExecuteChezmoi(args, false)

	if err != nil {
		// Print error to stderr and exit with failure code
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	// Exit with the same code that chezmoi returned
	os.Exit(exitCode)
}
