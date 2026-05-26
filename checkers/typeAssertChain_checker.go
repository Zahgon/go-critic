package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/checkers/internal/lintutil"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "typeAssertChain"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects repeated type assertions and suggests to replace them with type switch statement"
	info.Before = `
if x, ok := v.(T1); ok {
	// Code A, uses x.
} else if x, ok := v.(T2); ok {
	// Code B, uses x.
} else if x, ok := v.(T3); ok {
	// Code C, uses x.
}`
	info.After = `
switch x := v.(T1) {
case cond1:
	// Code A, uses x.
case cond2:
	// Code B, uses x.
default:
	// Code C, uses x.
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&typeAssertChainChecker{ctx: ctx}), nil
	})
}

type typeAssertChainChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	cause   *ast.IfStmt
	visited map[*ast.IfStmt]bool
	typeSet lintutil.AstSet
}

func (c *typeAssertChainChecker) EnterFunc(fn *ast.FuncDecl) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *typeAssertChainChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *typeAssertChainChecker) getTypeAssert(ifstmt *ast.IfStmt) *ast.TypeAssertExpr {
	_ = "STUB: not implemented"
	return nil
}

func (c *typeAssertChainChecker) checkIfStmt(stmt *ast.IfStmt, assertion *ast.TypeAssertExpr) {
	_ = "STUB: not implemented"
	return
}

func (c *typeAssertChainChecker) countTypeAssertions(stmt *ast.IfStmt, assertion *ast.TypeAssertExpr) int {
	_ = "STUB: not implemented"
	return 0
}

// Asserted type is duplicated.
// Type switch does not permit duplicate cases,
// so give up.

// Mixed type asserting chain.
// Can't be easily translated to a type switch.

func (c *typeAssertChainChecker) warn() { _ = "STUB: not implemented"; return }
