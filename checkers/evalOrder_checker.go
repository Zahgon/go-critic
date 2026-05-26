package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "evalOrder"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects unwanted dependencies on the evaluation order"
	info.Before = `return x, f(&x)`
	info.After = `
err := f(&x)
return x, err
`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&evalOrderChecker{ctx: ctx}), nil
	})
}

type evalOrderChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *evalOrderChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

// TODO(quasilyte): handle selector expressions like o.val in addition
// to bare identifiers.

// addrTake is &id now

// 1. Check if there is a call in form of id.method() where
// method takes id by a pointer.

// 2. Check that there is no call that uses &id as an argument.

func (c *evalOrderChecker) hasPtrRecv(fn *ast.Ident) bool { _ = "STUB: not implemented"; return false }

func (c *evalOrderChecker) warn(call *ast.CallExpr) { _ = "STUB: not implemented"; return }
