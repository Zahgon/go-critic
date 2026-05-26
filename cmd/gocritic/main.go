package main

import (
	"github.com/go-critic/go-critic/checkers"
)

var Version = "v0.0.0-SNAPSHOT"

func main() {
	err := checkers.InitEmbeddedRules()
	if err != nil {
		panic(err)
	}

	run(config{
		Name:    "go-critic",
		Version: Version,
	})
}

// config is used to parametrize the linter.
type config struct {
	Version string
	Name    string
}

// Run executes corresponding main after sub-command resolving.
// Does not return.
func run(cfg config) { _ = "STUB: not implemented"; return }
