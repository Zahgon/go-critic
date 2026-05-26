package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "ptrToRefParam"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Summary = "Detects input and output parameters that have a type of pointer to referential type"
	info.Before = `func f(m *map[string]int) (*chan *int)`
	info.After = `func f(m map[string]int) (chan *int)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&ptrToRefParamChecker{ctx: ctx}), nil
	})
}

type ptrToRefParamChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *ptrToRefParamChecker) VisitFuncDecl(fn *ast.FuncDecl) { _ = "STUB: not implemented"; return }

func (c *ptrToRefParamChecker) checkParams(params []*ast.Field) { _ = "STUB: not implemented"; return }

func (c *ptrToRefParamChecker) isRefType(x types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

// Handle underlying type only for interfaces.

func (c *ptrToRefParamChecker) warn(id *ast.Ident) { _ = "STUB: not implemented"; return }
