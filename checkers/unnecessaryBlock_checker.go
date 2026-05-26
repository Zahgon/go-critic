package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "unnecessaryBlock"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Summary = "Detects unnecessary braced statement blocks"
	info.Before = `
x := 1
{
	print(x)
}`
	info.After = `
x := 1
print(x)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmtList(&unnecessaryBlockChecker{ctx: ctx}), nil
	})
}

type unnecessaryBlockChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *unnecessaryBlockChecker) VisitStmtList(x ast.Node, statements []ast.Stmt) {
	_ = "STUB: not implemented"
	// Using StmtListVisitor instead of StmtVisitor makes it easier to avoid
	// false positives on IfStmt, RangeStmt, ForStmt and alike.
	// We only inspect BlockStmt inside statement lists, so this method is not
	// called for IfStmt itself, for example.
	return
}

func (c *unnecessaryBlockChecker) hasDefinitions(stmt *ast.BlockStmt) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *unnecessaryBlockChecker) warn(expr ast.Stmt) { _ = "STUB: not implemented"; return }
