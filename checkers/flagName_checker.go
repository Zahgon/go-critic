package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "flagName"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects suspicious flag names"
	info.Before = `b := flag.Bool(" foo ", false, "description")`
	info.After = `b := flag.Bool("foo", false, "description")`
	info.Note = "https://github.com/golang/go/issues/41792"

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&flagNameChecker{ctx: ctx}), nil
	})
}

type flagNameChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *flagNameChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *flagNameChecker) checkFlagName(call *ast.CallExpr, arg ast.Expr) {
	_ = "STUB: not implemented"
	return
}

// Non-constant name

func (c *flagNameChecker) warnEmpty(cause ast.Node) { _ = "STUB: not implemented"; return }

func (c *flagNameChecker) warnHyphenPrefix(cause ast.Node, name string) {
	_ = "STUB: not implemented"
	return
}

func (c *flagNameChecker) warnEq(cause ast.Node, name string) { _ = "STUB: not implemented"; return }

func (c *flagNameChecker) warnWhitespace(cause ast.Node, name string) {
	_ = "STUB: not implemented"
	return
}
