package checkers

import (
	"go/ast"
	"regexp"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "commentedOutCode"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects commented-out code inside function bodies"
	info.Params = linter.CheckerParams{
		"minLength": {
			Value: 15,
			Usage: "min length of the comment that triggers a warning",
		},
	}
	info.Before = `
// fmt.Println("Debugging hard")
foo(1, 2)`
	info.After = `foo(1, 2)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForLocalComment(&commentedOutCodeChecker{
			ctx:              ctx,
			notQuiteFuncCall: regexp.MustCompile(`\w+\s+\([^)]*\)\s*$`),
			minLength:        info.Params.Int("minLength"),
		}), nil
	})
}

type commentedOutCodeChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
	fn  *ast.FuncDecl

	notQuiteFuncCall *regexp.Regexp
	minLength        int
}

func (c *commentedOutCodeChecker) EnterFunc(fn *ast.FuncDecl) bool {
	_ = "STUB: not implemented"
	// Need to store current function inside checker context
	return false
}

func (c *commentedOutCodeChecker) VisitLocalComment(cg *ast.CommentGroup) {
	_ = "STUB: not implemented"
	// Collect text once
	return
}

// We do multiple heuristics to avoid false positives.
// Many things can be improved here.

// TODO comments with code are permitted.

// "http://" is interpreted as a label with comment.
// There are other protocols we might want to include.

// Clearly not a "selector expr" (mostly due to extra space)

// Some very short comment that can be skipped.
// Usually triggering on these results in false positive.
// Unless there is a very popular call like print/println.

// Almost looks like a commented-out function call,
// but there is a whitespace between function name and
// parameters list. Skip these to avoid false positives.

// Don't try to parse one-liner as block statement

// Add braces to make block statement from
// multiple statements.

// An example output comment can be one of the following:
//
//	Output: some output
//
// or
//
//	Output:
//	some output
//
// See https://go.dev/blog/examples
func (c *commentedOutCodeChecker) isExampleOutputComment(s string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *commentedOutCodeChecker) isPermittedStmt(stmt ast.Stmt) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *commentedOutCodeChecker) isPermittedExpr(x ast.Expr) bool {
	_ = "STUB: not implemented"
	// Permit anything except expressions that can be used
	// with complete result discarding.
	return false
}

// "<-" channel receive is not permitted.

func (c *commentedOutCodeChecker) warn(cause ast.Node) { _ = "STUB: not implemented"; return }
