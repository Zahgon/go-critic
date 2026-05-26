package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "methodExprCall"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects method expression call that can be replaced with a method call"
	info.Before = `f := foo{}
foo.bar(f)`
	info.After = `f := foo{}
f.bar()`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&methodExprCallChecker{ctx: ctx}), nil
	})
}

type methodExprCallChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *methodExprCallChecker) VisitExpr(x ast.Expr) { _ = "STUB: not implemented"; return }

func (c *methodExprCallChecker) warn(cause *ast.CallExpr, s *ast.SelectorExpr) {
	_ = "STUB: not implemented"
	return
}

// Remove "&" from the receiver (if any).
