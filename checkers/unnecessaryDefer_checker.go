package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "unnecessaryDefer"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects redundantly deferred calls"
	info.Before = `
func() {
	defer os.Remove(filename)
}`
	info.After = `
func() {
	os.Remove(filename)
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&unnecessaryDeferChecker{ctx: ctx}), nil
	})
}

type unnecessaryDeferChecker struct {
	astwalk.WalkHandler
	ctx    *linter.CheckerContext
	isFunc bool
}

// Visit implements the ast.Visitor. This visitor keeps track of the block
// statement belongs to a function or any other block. If the block is not a
// function and ends with a defer statement that should be OK since it's
// deferring the outer function.
func (c *unnecessaryDeferChecker) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func (c *unnecessaryDeferChecker) VisitFuncDecl(funcDecl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	// We always start as a function (*ast.FuncDecl.Body passed)
	return
}

func (c *unnecessaryDeferChecker) checkDeferBeforeReturn(funcDecl *ast.BlockStmt) {
	_ = "STUB: not implemented"
	// Check if we have an explicit return or if it's just the end of the scope.
	return
}

// If the block is a function and ending with return or if we have an
// explicit return in any other block we should warn about
// unnecessary defer.

func (c *unnecessaryDeferChecker) isTrivialReturn(ret *ast.ReturnStmt) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *unnecessaryDeferChecker) isConstExpr(e ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *unnecessaryDeferChecker) warn(deferStmt *ast.DeferStmt) { _ = "STUB: not implemented"; return }

// To avoid long and multi-line warning messages,
// collapse the function literals.
