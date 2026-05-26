package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "newDeref"
	info.Tags = []string{linter.StyleTag}
	info.Summary = "Detects immediate dereferencing of `new` expressions"
	info.Before = `x := *new(bool)`
	info.After = `x := false`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&newDerefChecker{ctx: ctx}), nil
	})
}

type newDerefChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *newDerefChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

// allow *new(T) if T is a type parameter, see #1272 for details

func (c *newDerefChecker) warn(cause, suggestion ast.Expr) { _ = "STUB: not implemented"; return }
