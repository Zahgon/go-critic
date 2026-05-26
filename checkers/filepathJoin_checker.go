package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "filepathJoin"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects problems in filepath.Join() function calls"
	info.Before = `filepath.Join("dir/", filename)`
	info.After = `filepath.Join("dir", filename)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&filepathJoinChecker{ctx: ctx}), nil
	})
}

type filepathJoinChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *filepathJoinChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *filepathJoinChecker) hasSeparator(v *ast.BasicLit) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *filepathJoinChecker) warnSeparator(sep ast.Expr) { _ = "STUB: not implemented"; return }
