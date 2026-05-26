package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "deferInLoop"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects loops inside functions that use defer"
	info.Before = `
for _, filename := range []string{"foo", "bar"} {
	 f, err := os.Open(filename)
	
	defer f.Close()
}
`
	info.After = `
func process(filename string) {
	 f, err := os.Open(filename)
	
	defer f.Close()
}
/* ... */
for _, filename := range []string{"foo", "bar"} {
	process(filename)
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&deferInLoopChecker{ctx: ctx}), nil
	})
}

type deferInLoopChecker struct {
	astwalk.WalkHandler
	ctx   *linter.CheckerContext
	inFor bool
}

func (c *deferInLoopChecker) VisitFuncDecl(fn *ast.FuncDecl) { _ = "STUB: not implemented"; return }

func (c deferInLoopChecker) traversalFunc(cur ast.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *deferInLoopChecker) warn(cause *ast.DeferStmt) { _ = "STUB: not implemented"; return }
