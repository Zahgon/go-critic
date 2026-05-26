package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "emptyFallthrough"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects fallthrough that can be avoided by using multi case values"
	info.Before = `switch kind {
case reflect.Int:
	fallthrough
case reflect.Int32:
	return Int
}`
	info.After = `switch kind {
case reflect.Int, reflect.Int32:
	return Int
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&emptyFallthroughChecker{ctx: ctx}), nil
	})
}

type emptyFallthroughChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *emptyFallthroughChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *emptyFallthroughChecker) warnDefault(cause ast.Node) { _ = "STUB: not implemented"; return }

func (c *emptyFallthroughChecker) warn(cause ast.Node) { _ = "STUB: not implemented"; return }
