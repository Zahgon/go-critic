package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "exitAfterDefer"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects calls to exit/fatal inside functions that use defer"
	info.Before = `
defer os.Remove(filename)
if bad {
	log.Fatalf("something bad happened")
}`
	info.After = `
defer os.Remove(filename)
if bad {
	log.Printf("something bad happened")
	return
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&exitAfterDeferChecker{ctx: ctx}), nil
	})
}

type exitAfterDeferChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *exitAfterDeferChecker) VisitFuncDecl(fn *ast.FuncDecl) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): handle goto and other kinds of flow that break
	// the algorithm below that expects the latter statement to be
	// executed after the ones that come before it.
	return
}

// If we found a defer statement in the function post traversal.
// and are looking at the Else branch during a pre traversal, stop seeking as it could be false positive.

// Don't recurse into local anonymous functions.

// See #995. We allow `defer os.Exit()` calls
// as it's harder to determine whether they're going
// to clutter anything without actually trying to
// simulate the defer stack + understanding the control flow.
// TODO: can we use CFG here?

func (c *exitAfterDeferChecker) warn(cause *ast.CallExpr, deferStmt *ast.DeferStmt) {
	_ = "STUB: not implemented"
	return
}

// To avoid long and multi-line warning messages,
// collapse the function literals.
