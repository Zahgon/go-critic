package linttest

import (
	"testing"
)

// Run executes integration tests.
func (cfg *IntegrationTest) Run(t *testing.T) { _ = "STUB: not implemented"; return }

func (cfg *IntegrationTest) runTest(t *testing.T, gocritic, gopath string) {
	_ = "STUB: not implemented"
	return
}

// If several tests re-use a single golden file,
// don't read it repeatedly, just re-use its contents.

// The format is:
//	runParams ... "|" goldenFile

// Read from a golden file or contents cache.

// Get the actual execution output.

// Copy parent env

// Override GOPATH.

// Error is prepended to the beginning.

// To get line-by-line diff, split is required.

func (cfg *IntegrationTest) buildLinter() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
