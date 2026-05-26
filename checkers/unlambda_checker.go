package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "unlambda"
	info.Tags = []string{linter.StyleTag}
	info.Summary = "Detects function literals that can be simplified"
	info.Before = `func(x int) int { return fn(x) }`
	info.After = `fn`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&unlambdaChecker{ctx: ctx}), nil
	})
}

type unlambdaChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *unlambdaChecker) VisitExpr(x ast.Expr) { _ = "STUB: not implemented"; return }

// Skip tricky cases; only handle simple calls

// See #762

// Permit only non-pointer struct method values.

// See #888 #1007

// Now check that all arguments match the parameters.

func (c *unlambdaChecker) warn(cause ast.Node, suggestion string) {
	_ = "STUB: not implemented"
	return
}

func (c *unlambdaChecker) lenArgs(args []ast.Expr) int { _ = "STUB: not implemented"; return 0 }

// Don't count function call. only args.
