package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "paramTypeCombine"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag}
	info.Summary = "Detects if function parameters could be combined by type and suggest the way to do it"
	info.Before = `func foo(a, b int, c, d int, e, f int, g int) {}`
	info.After = `func foo(a, b, c, d, e, f, g int) {}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&paramTypeCombineChecker{ctx: ctx}), nil
	})
}

type paramTypeCombineChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *paramTypeCombineChecker) EnterFunc(*ast.FuncDecl) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *paramTypeCombineChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

func (c *paramTypeCombineChecker) optimizeFuncType(f *ast.FuncType) *ast.FuncType {
	_ = "STUB: not implemented"
	return nil
}

func (c *paramTypeCombineChecker) optimizeParams(params *ast.FieldList) *ast.FieldList {
	_ = "STUB: not implemented"
	// To avoid false positives, skip unnamed param lists.
	//
	// We're using a property that Go only permits unnamed params
	// for the whole list, so it's enough to check whether any of
	// ast.Field have empty name list.
	return nil
}

func (c *paramTypeCombineChecker) warn(f1, f2 *ast.FuncType) { _ = "STUB: not implemented"; return }

func (c *paramTypeCombineChecker) paramsAreMultiLine(params *ast.FieldList) bool {
	_ = "STUB: not implemented"
	return false
}
