package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "boolExprSimplify"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects bool expressions that can be simplified"
	info.Before = `
a := !(elapsed >= expectElapsedMin)
b := !(x) == !(y)`
	info.After = `
a := elapsed < expectElapsedMin
b := (x) == (y)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&boolExprSimplifyChecker{ctx: ctx}), nil
	})
}

type boolExprSimplifyChecker struct {
	astwalk.WalkHandler
	ctx       *linter.CheckerContext
	hasFloats bool
}

func (c *boolExprSimplifyChecker) VisitExpr(x ast.Expr) { _ = "STUB: not implemented"; return }

// Throw away non-bool expressions and avoid redundant
// AST copying below.

// We'll loose all types info after a copy,
// this is why we record valuable info before doing it.

func (c *boolExprSimplifyChecker) simplifyBool(x ast.Expr) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func (c *boolExprSimplifyChecker) doubleNegation(cur *astutil.Cursor) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *boolExprSimplifyChecker) negatedEquals(cur *astutil.Cursor) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *boolExprSimplifyChecker) invertComparison(cur *astutil.Cursor) bool {
	_ = "STUB: not implemented"
	// See #673
	return false
}

// Replace operator to its negated form.

func (c *boolExprSimplifyChecker) isSafe(x ast.Expr) bool { _ = "STUB: not implemented"; return false }

func (c *boolExprSimplifyChecker) combineChecks(cur *astutil.Cursor) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *boolExprSimplifyChecker) removeIncDec(cur *astutil.Cursor) bool {
	_ = "STUB: not implemented"
	return false
}

// `x > y-1` => `x >= y`
// `x+1 > y` => `x >= y`

// `x >= y+1` => `x > y`
// `x-1 >= y` => `x > y`

// `x < y+1` => `x <= y`
// `x-1 < y` => `x <= y`

// `x <= y-1` => `x < y`
// `x+1 <= y` => `x < y`

func (c *boolExprSimplifyChecker) foldRanges(cur *astutil.Cursor) bool {
	_ = "STUB: not implemented"
	// See #848
	return false
}

// `x > c && x < c+2` => `x == c+1`

// `x >= c && x < c+1` => `x == c`

// `x > c && x <= c+1` => `x == c+1`

// `x >= c && x <= c` => `x == c`

// `x < c || x > c` => `x != c`

// `x <= c || x > c+1` => `x != c+1`

// `x < c || x >= c+1` => `x != c`

// `x <= c || x >= c+2` => `x != c+1`

func (c *boolExprSimplifyChecker) int64val(x ast.Expr) (int64, bool) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): if we had types info, we could use TypesInfo.Types[x].Value,
	// but since copying erases leaves us without it, only basic literals are handled.
	return 0, false
}

func (c *boolExprSimplifyChecker) warn(cause, suggestion ast.Expr) {
	_ = "STUB: not implemented"
	return
}
