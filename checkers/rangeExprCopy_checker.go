package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "rangeExprCopy"
	info.Tags = []string{linter.PerformanceTag}
	info.Params = linter.CheckerParams{
		"sizeThreshold": {
			Value: 512,
			Usage: "size in bytes that makes the warning trigger",
		},
		"skipTestFuncs": {
			Value: true,
			Usage: "whether to check test functions",
		},
	}
	info.Summary = "Detects expensive copies of `for` loop range expressions"
	info.Details = "Suggests to use pointer to array to avoid the copy using `&` on range expression."
	info.Before = `
var xs [2048]byte
for _, x := range xs { // Copies 2048 bytes
	// Loop body.
}`
	info.After = `
var xs [2048]byte
for _, x := range &xs { // No copy
	// Loop body.
}`
	info.Note = "See Go issue for details: https://github.com/golang/go/issues/15812."

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &rangeExprCopyChecker{ctx: ctx}
		c.sizeThreshold = int64(info.Params.Int("sizeThreshold"))
		c.skipTestFuncs = info.Params.Bool("skipTestFuncs")
		return astwalk.WalkerForStmt(c), nil
	})
}

type rangeExprCopyChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	sizeThreshold int64
	skipTestFuncs bool
}

func (c *rangeExprCopyChecker) EnterFunc(fn *ast.FuncDecl) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *rangeExprCopyChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *rangeExprCopyChecker) warn(rng *ast.RangeStmt, size int64) {
	_ = "STUB: not implemented"
	return
}
