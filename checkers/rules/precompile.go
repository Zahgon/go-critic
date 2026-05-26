//go:build generate

package main

import (
	"log"
)

// This program generates a loadable IR for ruleguard
// so we don't have to load the rules from AST and typecheck
// them every time.

func main() {
	log.SetFlags(0)
	if err := precompile(); err != nil {
		log.Printf("error: %v", err)
	}
}

func precompile() error { _ = "STUB: not implemented"; return nil }
