package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "appendCombine"
	info.Tags = []string{linter.PerformanceTag}
	info.Summary = "Detects `append` chains to the same slice that can be done in a single `append` call"
	info.Before = `
xs = append(xs, 1)
xs = append(xs, 2)`
	info.After = `xs = append(xs, 1, 2)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmtList(&appendCombineChecker{ctx: ctx}), nil
	})
}

type appendCombineChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *appendCombineChecker) VisitStmtList(_ ast.Node, list []ast.Stmt) {
	_ = "STUB: not implemented"
	// First append
	return
}

// Slice being appended to
// How much appends in a row we've seen

// Break the chain.
// If enough appends are in chain, print warning.

// First append in a chain.

// Required for printing chains that consist of trailing
// statements from the list.

func (c *appendCombineChecker) matchAppend(stmt ast.Stmt, slice ast.Expr) *ast.CallExpr {
	_ = "STUB: not implemented"
	// Seeking for:
	//	slice = append(slice, xs...)
	// xs are 0-N append arguments, but not variadic argument,
	// because it makes append combining impossible.
	return nil
}

// Check that current append slice match previous append slice.
// Otherwise we should break the chain.

func (c *appendCombineChecker) warn(cause ast.Node, chain int) { _ = "STUB: not implemented"; return }
