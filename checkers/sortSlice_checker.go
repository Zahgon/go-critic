package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "sortSlice"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious sort.Slice calls"
	info.Before = `sort.Slice(xs, func(i, j) bool { return keys[i] < keys[j] })`
	info.After = `sort.Slice(kv, func(i, j) bool { return kv[i].key < kv[j].key })`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&sortSliceChecker{ctx: ctx}), nil
	})
}

type sortSliceChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *sortSliceChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

// OK.

// Don't check unpredictable slice values

// Both cmp.X and cmp.Y are expected to be some expressions
// over the `slice` expression. In the simplest case,
// it's a `slice[i] <op> slice[j]`.

// This one is more about the style, but can reveal potential issue
// or misprint in sorting condition.
// We give a warn if X contains indexing with `i` index and Y
// contains indexing with `j`.

func (c *sortSliceChecker) paramIdents(e *ast.FuncType) (ivar, jvar *ast.Ident) {
	_ = "STUB: not implemented"
	// Covers both `i, j int` and `i int, j int`.
	return nil, nil
}

func (c *sortSliceChecker) unwrapSlice(e ast.Expr) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (c *sortSliceChecker) containsIndex(e, index ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *sortSliceChecker) containsSlice(e, slice ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *sortSliceChecker) warnSlice(cause ast.Node, slice ast.Expr) {
	_ = "STUB: not implemented"
	return
}

func (c *sortSliceChecker) warnIndex(cause ast.Node, ivar, jvar *ast.Ident) {
	_ = "STUB: not implemented"
	return
}
