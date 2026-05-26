package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "tooManyResultsChecker"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"maxResults": {
			Value: 5,
			Usage: "maximum number of results",
		},
	}
	info.Summary = "Detects function with too many results"
	info.Before = `func fn() (a, b, c, d float32, _ int, _ bool)`
	info.After = `func fn() (resultStruct, bool)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := astwalk.WalkerForFuncDecl(&tooManyResultsChecker{
			ctx:       ctx,
			maxParams: info.Params.Int("maxResults"),
		})
		return c, nil
	})
}

type tooManyResultsChecker struct {
	astwalk.WalkHandler
	ctx       *linter.CheckerContext
	maxParams int
}

func (c *tooManyResultsChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

func (c *tooManyResultsChecker) warn(n ast.Node) { _ = "STUB: not implemented"; return }
