package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "truncateCmp"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"skipArchDependent": {
			Value: true,
			Usage: "whether to skip int/uint/uintptr types",
		},
	}
	info.Summary = "Detects potential truncation issues when comparing ints of different sizes"
	info.Before = `
func f(x int32, y int16) bool {
  return int16(x) < y
}`
	info.After = `
func f(x int32, int16) bool {
  return x < int32(y)
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &truncateCmpChecker{ctx: ctx}
		c.skipArchDependent = info.Params.Bool("skipArchDependent")
		return astwalk.WalkerForExpr(c), nil
	})
}

type truncateCmpChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	skipArchDependent bool
}

func (c *truncateCmpChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

// Don't bother about untyped consts

func (c *truncateCmpChecker) isTruncCast(x ast.Expr) bool { _ = "STUB: not implemented"; return false }

func (c *truncateCmpChecker) checkCmp(cmpX, cmpY ast.Expr) {
	_ = "STUB: not implemented"
	// Check if we have a cast to a type that can truncate.
	return
}

// Just in case of the shadowed builtin

// Check that both x and y are signed or unsigned int-typed.

func (c *truncateCmpChecker) warn(cause ast.Expr, xsize, ysize int64, suggest string) {
	_ = "STUB: not implemented"
	return
}
