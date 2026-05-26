package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "typeUnparen"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag}
	info.Summary = "Detects unneeded parenthesis inside type expressions and suggests to remove them"
	info.Before = `type foo [](func([](func())))`
	info.After = `type foo []func([]func())`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForTypeExpr(&typeUnparenChecker{ctx: ctx}, ctx.TypesInfo), nil
	})
}

type typeUnparenChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *typeUnparenChecker) VisitTypeExpr(e ast.Expr) { _ = "STUB: not implemented"; return }

// Only nested fields are to be reported.

func (c *typeUnparenChecker) checkType(e ast.Expr) { _ = "STUB: not implemented"; return }

func (c *typeUnparenChecker) removeRedundantParens(e ast.Expr) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (c *typeUnparenChecker) warn(cause, noParens ast.Expr) { _ = "STUB: not implemented"; return }
