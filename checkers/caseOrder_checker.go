package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "caseOrder"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects erroneous case order inside switch statements"
	info.Before = `
switch x.(type) {
case ast.Expr:
	fmt.Println("expr")
case *ast.BasicLit:
	fmt.Println("basic lit") // Never executed
}`
	info.After = `
switch x.(type) {
case *ast.BasicLit:
	fmt.Println("basic lit") // Now reachable
case ast.Expr:
	fmt.Println("expr")
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&caseOrderChecker{ctx: ctx}), nil
	})
}

type caseOrderChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *caseOrderChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *caseOrderChecker) checkTypeSwitch(s *ast.TypeSwitchStmt) {
	_ = "STUB: not implemented"
	return
}

// Interfaces seen so far

func (c *caseOrderChecker) warnTypeSwitch(cause, concrete, iface ast.Node) {
	_ = "STUB: not implemented"
	return
}

func (c *caseOrderChecker) warnUnknownType(cause, concrete ast.Node) {
	_ = "STUB: not implemented"
	return
}

func (c *caseOrderChecker) checkSwitch(_ *ast.SwitchStmt) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): can handle expression cases that overlap.
	// Cases that have narrower value range should go before wider ones.
	return
}
