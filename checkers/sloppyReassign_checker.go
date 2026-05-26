package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "sloppyReassign"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious/confusing re-assignments"
	info.Before = `if err = f(); err != nil { return err }`
	info.After = `if err := f(); err != nil { return err }`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&sloppyReassignChecker{ctx: ctx}), nil
	})
}

type sloppyReassignChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *sloppyReassignChecker) VisitStmt(stmt ast.Stmt) {
	_ = "STUB: not implemented"
	// Right now only check assignments in if statements init.
	return
}

// TODO(quasilyte): is handling of multi-value assignments worthwhile?

// TODO(quasilyte): handle not only the simplest, return-only case.

// Variable that is being re-assigned.

// TODO(quasilyte): handle not only nil comparisons.

func (c *sloppyReassignChecker) warnAssignToDefine(assign *ast.AssignStmt, name string) {
	_ = "STUB: not implemented"
	return
}
