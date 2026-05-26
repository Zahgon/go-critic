package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "initClause"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Summary = "Detects non-assignment statements inside if/switch init clause"
	info.Before = `if sideEffect(); cond {
}`
	info.After = `sideEffect()
if cond {
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&initClauseChecker{ctx: ctx}), nil
	})
}

type initClauseChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *initClauseChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *initClauseChecker) getInitClause(x ast.Stmt) ast.Stmt {
	_ = "STUB: not implemented"
	return *new(ast.Stmt)
}

func (c *initClauseChecker) warn(stmt, clause ast.Stmt) { _ = "STUB: not implemented"; return }
