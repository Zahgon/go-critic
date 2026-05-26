package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "sloppyTypeAssert"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects redundant type assertions"
	info.Before = `
func f(r io.Reader) interface{} {
	return r.(interface{})
}
`
	info.After = `
func f(r io.Reader) interface{} {
	return r
}
`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&sloppyTypeAssertChecker{ctx: ctx}), nil
	})
}

type sloppyTypeAssertChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *sloppyTypeAssertChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *sloppyTypeAssertChecker) warnIdentical(cause ast.Expr) { _ = "STUB: not implemented"; return }
