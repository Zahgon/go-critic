package checkers

import (
	"go/ast"
	"go/token"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "badCond"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects suspicious condition expressions"
	info.Before = `
for i := 0; i > n; i++ {
	xs[i] = 0
}`
	info.After = `
for i := 0; i < n; i++ {
	xs[i] = 0
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&badCondChecker{ctx: ctx}), nil
	})
}

type badCondChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *badCondChecker) VisitFuncDecl(decl *ast.FuncDecl) { _ = "STUB: not implemented"; return }

func (c *badCondChecker) checkExpr(expr ast.Expr) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): recognize more patterns.
	return
}

// Notes:
// `x != a || x != b` handled by go vet.

// Pattern 1.
// `x < a && x > b`; Where `a` is less than `b`.

// Pattern 2.
// `x == a && x == b`
//
// Valid when `b == a` is intended, but still reported.
// We can disable "just suspicious" warnings by default
// is users are upset with the current behavior.

func (c *badCondChecker) equalToBoth(lhs, rhs *ast.BinaryExpr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *badCondChecker) lessAndGreater(lhs, rhs *ast.BinaryExpr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *badCondChecker) checkForStmt(stmt *ast.ForStmt) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): handle other kinds of bad conditionals.
	return
}

func (c *badCondChecker) warnForStmt(cause ast.Node, op token.Token, cond *ast.BinaryExpr) {
	_ = "STUB: not implemented"
	return
}

func (c *badCondChecker) warnCond(cond *ast.BinaryExpr, tag string) {
	_ = "STUB: not implemented"
	return
}
