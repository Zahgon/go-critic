package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "rangeAppendAll"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects append all its data while range it"
	info.Before = `for _, n := range ns {
	...
		rs = append(rs, ns...) // append all slice data
	}
}`
	info.After = `for _, n := range ns {
	...
		rs = append(rs, n)
	}
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &rangeAppendAllChecker{ctx: ctx}
		return astwalk.WalkerForStmt(c), nil
	})
}

type rangeAppendAllChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *rangeAppendAllChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *rangeAppendAllChecker) getValidAppendFrom(expr ast.Node) *ast.Ident {
	_ = "STUB: not implemented"
	return nil
}

func (c *rangeAppendAllChecker) isSliceLiteral(arg ast.Expr) bool {
	_ = "STUB: not implemented"
	return false

	// []T{}, []T{n}
}

// []T(nil)

func (c *rangeAppendAllChecker) warn(appendFrom *ast.Ident) { _ = "STUB: not implemented"; return }
