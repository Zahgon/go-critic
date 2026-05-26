package checkers

import (
	"go/ast"
	"go/build"
	"go/token"
	"sync"

	"github.com/go-critic/go-critic/linter"

	"github.com/quasilyte/go-ruleguard/ruleguard"
)

//go:generate go run ./rules/precompile.go -rules ./rules/rules.go -o ./rulesdata/rulesdata.go

// cachedEngine holds a pre-initialized ruleguard engine for a specific rule group.
// The engine is created once and reused for all checker instances.
type cachedEngine struct {
	engine *ruleguard.Engine
	once   sync.Once
	err    error

	// Configuration needed to create the engine
	fset         *token.FileSet
	buildContext *build.Context
	groupName    string
	debug        bool
}

func (ce *cachedEngine) get() (*ruleguard.Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InitEmbeddedRules() error { _ = "STUB: not implemented"; return nil }

// First we create an Engine to parse all rules.
// We need it to get the structured info about our rules
// that will be used to generate checkers.
// We introduce an extra scope in hope that rootEngine
// will be garbage-collected after we don't need it.
// LoadedGroups() returns a slice copy and that's all what we need.

// For every rules group we create a cached engine holder.
// The engine will be created lazily on first use and then reused.

// Create a cached engine for this rule group

type embeddedRuleguardChecker struct {
	ctx    *linter.CheckerContext
	engine *ruleguard.Engine
}

func (c *embeddedRuleguardChecker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
