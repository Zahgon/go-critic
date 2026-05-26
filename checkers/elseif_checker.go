package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "elseif"
	info.Tags = []string{linter.StyleTag}
	info.Params = linter.CheckerParams{
		"skipBalanced": {
			Value: true,
			Usage: "whether to skip balanced if-else pairs",
		},
	}
	info.Summary = "Detects else with nested if statement that can be replaced with else-if"
	info.Before = `
if cond1 {
} else {
	if x := cond2; x {
	}
}`
	info.After = `
if cond1 {
} else if x := cond2; x {
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &elseifChecker{ctx: ctx}
		c.skipBalanced = info.Params.Bool("skipBalanced")
		return astwalk.WalkerForStmt(c), nil
	})
}

type elseifChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	skipBalanced bool
}

func (c *elseifChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

// Configured to skip balanced statements

func (c *elseifChecker) warn(cause ast.Node) { _ = "STUB: not implemented"; return }
