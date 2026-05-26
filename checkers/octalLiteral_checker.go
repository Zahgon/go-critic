package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "octalLiteral"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag, linter.OpinionatedTag}
	info.Summary = "Detects old-style octal literals"
	info.Before = `foo(02)`
	info.After = `foo(0o2)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&octalLiteralChecker{ctx: ctx}), nil
	})
}

type octalLiteralChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *octalLiteralChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *octalLiteralChecker) warn(lit *ast.BasicLit) { _ = "STUB: not implemented"; return }
