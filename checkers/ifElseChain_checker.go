package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "ifElseChain"
	info.Tags = []string{linter.StyleTag}
	info.Params = linter.CheckerParams{
		"minThreshold": {
			Value: 2,
			Usage: "min number of if-else blocks that makes the warning trigger",
		},
	}
	info.Summary = "Detects repeated if-else statements and suggests to replace them with switch statement"
	info.Before = `
if cond1 {
	// Code A.
} else if cond2 {
	// Code B.
} else {
	// Code C.
}`
	info.After = `
switch {
case cond1:
	// Code A.
case cond2:
	// Code B.
default:
	// Code C.
}`
	info.Note = `
Permits single else or else-if; repeated else-if or else + else-if
will trigger suggestion to use switch statement.
See [EffectiveGo#switch](https://golang.org/doc/effective_go.html#switch).`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&ifElseChainChecker{
			ctx:          ctx,
			minThreshold: info.Params.Int("minThreshold"),
		}), nil
	})
}

type ifElseChainChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	cause   *ast.IfStmt
	visited map[*ast.IfStmt]bool

	minThreshold int
}

func (c *ifElseChainChecker) EnterFunc(fn *ast.FuncDecl) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ifElseChainChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

func (c *ifElseChainChecker) checkIfStmt(stmt *ast.IfStmt) { _ = "STUB: not implemented"; return }

func (c *ifElseChainChecker) countIfelseLen(stmt *ast.IfStmt) int {
	_ = "STUB: not implemented"
	return 0
}

// Give up

// Else if.

// Else branch.

// No else or else if.

func (c *ifElseChainChecker) warn() { _ = "STUB: not implemented"; return }
