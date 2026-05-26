package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "weakCond"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects conditions that are unsafe due to not being exhaustive"
	info.Before = `xs != nil && xs[0] != nil`
	info.After = `len(xs) != 0 && xs[0] != nil`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&weakCondChecker{ctx: ctx}), nil
	})
}

type weakCondChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *weakCondChecker) VisitExpr(expr ast.Expr) {
	_ = "STUB: not implemented"
	// TODO(Quasilyte): more patterns.
	// TODO(Quasilyte): analyze and fix false positives.
	return
}

// Pattern 1.
// `x != nil && usageOf(x[i])`
// Pattern 2.
// `x == nil || usageOf(x[i])`

// lhs is `x <op> nil`

// isIndexed reports whether x is indexed inside given expr tree.
func (c *weakCondChecker) isIndexed(tree, x ast.Expr) bool { _ = "STUB: not implemented"; return false }

func (c *weakCondChecker) warn(cause ast.Node, suggest string) { _ = "STUB: not implemented"; return }
