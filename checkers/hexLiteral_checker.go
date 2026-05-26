package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "hexLiteral"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects hex literals that have mixed case letter digits"
	info.Before = `
x := 0X12
y := 0xfF`
	info.After = `
x := 0x12
// (A)
y := 0xff
// (B)
y := 0xFF`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&hexLiteralChecker{ctx: ctx}), nil
	})
}

type hexLiteralChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *hexLiteralChecker) warn0X(lit *ast.BasicLit) { _ = "STUB: not implemented"; return }

func (c *hexLiteralChecker) warnMixedDigits(lit *ast.BasicLit) { _ = "STUB: not implemented"; return }

func (c *hexLiteralChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }
