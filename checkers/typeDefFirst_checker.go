package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "typeDefFirst"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects method declarations preceding the type definition itself"
	info.Before = `
func (r rec) Method() {}
type rec struct{}
`
	info.After = `
type rec struct{}
func (r rec) Method() {}
`
	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &typeDefFirstChecker{
			ctx: ctx,
		}, nil
	})
}

type typeDefFirstChecker struct {
	astwalk.WalkHandler
	ctx          *linter.CheckerContext
	trackedTypes map[string]bool
}

func (c *typeDefFirstChecker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

func (c *typeDefFirstChecker) walkDecl(decl ast.Decl) { _ = "STUB: not implemented"; return }

func (c *typeDefFirstChecker) receiverType(e ast.Expr) string { _ = "STUB: not implemented"; return "" }

func (c *typeDefFirstChecker) warn(cause ast.Node, typeName string) {
	_ = "STUB: not implemented"
	return
}
