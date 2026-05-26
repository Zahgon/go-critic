package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "nilValReturn"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects return statements those results evaluate to nil"
	info.Before = `
if err == nil {
	return err
}`
	info.After = `
// (A) - return nil explicitly
if err == nil {
	return nil
}
// (B) - typo in "==", change to "!="
if err != nil {
	return err
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&nilValReturnChecker{ctx: ctx}), nil
	})
}

type nilValReturnChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *nilValReturnChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *nilValReturnChecker) warn(cause, val ast.Node) { _ = "STUB: not implemented"; return }
