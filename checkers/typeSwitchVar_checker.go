package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "typeSwitchVar"
	info.Tags = []string{linter.StyleTag}
	info.Summary = "Detects type switches that can benefit from type guard clause with variable"
	info.Before = `
switch v.(type) {
case int:
	return v.(int)
case point:
	return v.(point).x + v.(point).y
default:
	return 0
}`
	info.After = `
switch v := v.(type) {
case int:
	return v
case point:
	return v.x + v.y
default:
	return 0
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&typeSwitchVarChecker{ctx: ctx}), nil
	})
}

type typeSwitchVarChecker struct {
	astwalk.WalkHandler
	ctx   *linter.CheckerContext
	count int
}

func (c *typeSwitchVarChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *typeSwitchVarChecker) checkTypeSwitch(root *ast.TypeSwitchStmt) {
	_ = "STUB: not implemented"
	return
}

// Already with type guard

// Must be a *ast.ExprStmt then.

// Give up: can't handle shadowing without object

// Multiple types in a list mean that assert.X will have
// a type of interface{} inside clause body.
// We are looking for precise type case.

// Create artificial node just for matching.

func (c *typeSwitchVarChecker) warn(n ast.Node) { _ = "STUB: not implemented"; return }
