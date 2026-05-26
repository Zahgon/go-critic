package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "appendAssign"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects suspicious append result assignments"
	info.Before = `
p.positives = append(p.negatives, x)
p.negatives = append(p.negatives, y)`
	info.After = `
p.positives = append(p.positives, x)
p.negatives = append(p.negatives, y)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&appendAssignChecker{ctx: ctx}), nil
	})
}

type appendAssignChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *appendAssignChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *appendAssignChecker) checkAppend(x ast.Expr, call *ast.CallExpr) {
	_ = "STUB: not implemented"
	return
}

// Try to detect `xs = append(ys, xs...)` idiom.

// Don't check assignments to blank ident

// Most likely `m[k] = append(x, ...)`
// pattern, where x was retrieved by m[k] before.
//
// TODO: it's possible to record such map/slice reads
// and check whether it was done before this call.
// But for now, treat it like x belongs to m[k].

// Arrays are frequently used as scratch storages.

func (c *appendAssignChecker) matchSlices(cause ast.Node, x, y ast.Expr) {
	_ = "STUB: not implemented"
	return
}

func (c *appendAssignChecker) warn(cause ast.Node) { _ = "STUB: not implemented"; return }
