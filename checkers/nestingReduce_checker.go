package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "nestingReduce"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"bodyWidth": {
			Value: 5,
			Usage: "min number of statements inside a branch to trigger a warning",
		},
	}
	info.Summary = "Finds where nesting level could be reduced"
	info.Before = `
for _, v := range a {
	if v.Bool {
		body()
	}
}`
	info.After = `
for _, v := range a {
	if !v.Bool {
		continue
	}
	body()
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &nestingReduceChecker{ctx: ctx}
		c.bodyWidth = info.Params.Int("bodyWidth")
		return astwalk.WalkerForStmt(c), nil
	})
}

type nestingReduceChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	bodyWidth int
}

func (c *nestingReduceChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *nestingReduceChecker) checkLoopBody(body []ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *nestingReduceChecker) warnLoop(cause ast.Node) { _ = "STUB: not implemented"; return }
