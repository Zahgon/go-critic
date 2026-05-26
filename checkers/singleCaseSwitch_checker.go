package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "singleCaseSwitch"
	info.Tags = []string{linter.StyleTag}
	info.Summary = "Detects switch statements that could be better written as if statement"
	info.Before = `
switch x := x.(type) {
case int:
	body()
}`
	info.After = `
if x, ok := x.(int); ok {
	body()
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&singleCaseSwitchChecker{ctx: ctx}), nil
	})
}

type singleCaseSwitchChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *singleCaseSwitchChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *singleCaseSwitchChecker) checkSwitchStmt(stmt ast.Stmt, body *ast.BlockStmt) {
	_ = "STUB: not implemented"
	return
}

func (c *singleCaseSwitchChecker) hasBreak(stmt ast.Stmt) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *singleCaseSwitchChecker) warn(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *singleCaseSwitchChecker) warnDefault(stmt ast.Stmt) { _ = "STUB: not implemented"; return }
