package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "unlabelStmt"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects redundant statement labels"
	info.Before = `
derp:
for x := range xs {
	if x == 0 {
		break derp
	}
}`
	info.After = `
for x := range xs {
	if x == 0 {
		break
	}
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&unlabelStmtChecker{ctx: ctx}), nil
	})
}

type unlabelStmtChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *unlabelStmtChecker) EnterFunc(fn *ast.FuncDecl) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO(quasilyte): should not do additional traversal here.
// For now, skip all functions that contain goto statement.

func (c *unlabelStmtChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

// We have a labeled statement from that have labeled continue/break.
// This is an invariant, since unused label is a compile-time error
// and we're currently skipping functions containing goto.
//
// Also note that Go labels are function-scoped and there
// can be no re-definitions. This means that we don't
// need to care about label shadowing or things like that.
//
// The task is to find cases where labeled branch (continue/break)
// is redundant and can be re-written, decreasing the label usages
// and potentially leading to its redundancy,
// or finding the redundant labels right away.

// Simplest case that can prove that label is redundant.
//
// If labeled branch is somewhere inside the statement block itself
// and none of the nested break'able statements refer to that label,
// the label can be removed.

// Only for loops: if last stmt in list is a loop
// that contains labeled "continue" to the outer loop label,
// it can be refactored to use "break" instead.
// Exceptions: select statements with a labeled "continue" are ignored.

// isLoop reports whether n is a loop of some kind.
// In other words, it tells whether n body can contain "continue"
// associated with n.
func (c *unlabelStmtChecker) isLoop(n ast.Node) bool { _ = "STUB: not implemented"; return false }

// canBreakFrom reports whether it is possible to "break" or "continue" from n body.
func (c *unlabelStmtChecker) canBreakFrom(n ast.Node) bool { _ = "STUB: not implemented"; return false }

// blockStmtOf returns body of specified node.
//
// TODO(quasilyte): handle other statements and see if it can be useful
// in other checkers.
func (c *unlabelStmtChecker) blockStmtOf(n ast.Node) *ast.BlockStmt {
	_ = "STUB: not implemented"
	return nil
}

// usesLabel reports whether n contains a usage of label.
func (c *unlabelStmtChecker) usesLabel(n *ast.BlockStmt, label string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *unlabelStmtChecker) warnRedundant(cause *ast.LabeledStmt) {
	_ = "STUB: not implemented"
	return
}

func (c *unlabelStmtChecker) warnLabeledContinue(cause ast.Node, label string) {
	_ = "STUB: not implemented"
	return
}
