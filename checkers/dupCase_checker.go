package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/checkers/internal/lintutil"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "dupCase"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects duplicated case clauses inside switch or select statements"
	info.Before = `
switch x {
case ys[0], ys[1], ys[2], ys[0], ys[4]:
}`
	info.After = `
switch x {
case ys[0], ys[1], ys[2], ys[3], ys[4]:
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&dupCaseChecker{ctx: ctx}), nil
	})
}

type dupCaseChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	astSet lintutil.AstSet
}

func (c *dupCaseChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *dupCaseChecker) checkSwitch(stmt *ast.SwitchStmt) { _ = "STUB: not implemented"; return }

func (c *dupCaseChecker) checkSelect(stmt *ast.SelectStmt) { _ = "STUB: not implemented"; return }

func (c *dupCaseChecker) warn(cause ast.Node) { _ = "STUB: not implemented"; return }
