package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "unnamedResult"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"checkExported": {
			Value: false,
			Usage: "whether to check exported functions",
		},
	}
	info.Summary = "Detects unnamed results that may benefit from names"
	info.Before = `func f() (float64, float64)`
	info.After = `func f() (x, y float64)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &unnamedResultChecker{ctx: ctx}
		c.checkExported = info.Params.Bool("checkExported")
		return astwalk.WalkerForFuncDecl(c), nil
	})
}

type unnamedResultChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	checkExported bool
}

func (c *unnamedResultChecker) VisitFuncDecl(decl *ast.FuncDecl) { _ = "STUB: not implemented"; return }

// Function has no results

// Skip named results

// Main difference with case of len=2 is that we permit any
// typ1 as long as second type is either error or bool.

func (c *unnamedResultChecker) typeName(typ types.Type) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *unnamedResultChecker) warn(n ast.Node) { _ = "STUB: not implemented"; return }
