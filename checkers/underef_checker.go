package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "underef"
	info.Tags = []string{linter.StyleTag}
	info.Params = linter.CheckerParams{
		"skipRecvDeref": {
			Value: true,
			Usage: "whether to skip (*x).method() calls where x is a pointer receiver",
		},
	}
	info.Summary = "Detects dereference expressions that can be omitted"
	info.Before = `
(*k).field = 5
v := (*a)[5] // only if a is array`
	info.After = `
k.field = 5
v := a[5]`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &underefChecker{ctx: ctx}
		c.skipRecvDeref = info.Params.Bool("skipRecvDeref")
		return astwalk.WalkerForExpr(c), nil
	})
}

type underefChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	skipRecvDeref bool
}

func (c *underefChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *underefChecker) isPtrRecvMethodCall(fn *ast.Ident) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *underefChecker) underef(x *ast.ParenExpr) ast.Expr {
	_ = "STUB: not implemented"
	// If there is only 1 deref, can remove parenthesis,
	// otherwise can remove StarExpr only.
	return *new(ast.Expr)
}

func (c *underefChecker) warnSelect(expr *ast.SelectorExpr) {
	_ = "STUB: not implemented"
	// TODO: add () to function output.
	return
}

func (c *underefChecker) warnArray(expr *ast.IndexExpr) { _ = "STUB: not implemented"; return }

// checkStarExpr checks if ast.StarExpr could be simplified.
func (c *underefChecker) checkStarExpr(expr *ast.StarExpr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *underefChecker) checkArray(expr *ast.StarExpr) bool {
	_ = "STUB: not implemented"
	return false
}
