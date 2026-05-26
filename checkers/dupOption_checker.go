package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "dupOption"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects duplicated option function arguments in variadic function calls"
	info.Before = `doSomething(name,
		withWidth(w),
		withHeight(h),
		withWidth(w),
)`
	info.After = `doSomething(name,
		withWidth(w),
		withHeight(h),
)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &dupOptionChecker{ctx: ctx}
		return astwalk.WalkerForExpr(c), nil
	})
}

type dupOptionChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *dupOptionChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *dupOptionChecker) getVariadicArgs(call *ast.CallExpr) ([]ast.Expr, types.Type) {
	_ = "STUB: not implemented"
	return nil, *new(types.Type)
}

// skip for someFunc(a, b ...)

func (c *dupOptionChecker) isOptionType(typeInfo types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *dupOptionChecker) findDupArgs(args []ast.Expr) []ast.Expr {
	_ = "STUB: not implemented"
	return nil
}

func (c *dupOptionChecker) warn(arg ast.Node) { _ = "STUB: not implemented"; return }
