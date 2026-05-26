package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/checkers/internal/lintutil"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "mapKey"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects suspicious map literal keys"
	info.Before = `
_ = map[string]int{
	"foo": 1,
	"bar ": 2,
}`
	info.After = `
_ = map[string]int{
	"foo": 1,
	"bar": 2,
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&mapKeyChecker{ctx: ctx}), nil
	})
}

type mapKeyChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	astSet lintutil.AstSet
}

func (c *mapKeyChecker) VisitExpr(expr ast.Expr) { _ = "STUB: not implemented"; return }

func (c *mapKeyChecker) checkDuplicates(lit *ast.CompositeLit) { _ = "STUB: not implemented"; return }

// Basic lits are handled by the compiler.

func (c *mapKeyChecker) checkWhitespace(lit *ast.CompositeLit) { _ = "STUB: not implemented"; return }

// s is unquoted string literal value.

// Already seen something with a whitespace.
// More than one entry => not suspicious.

// If space is used as a key, maybe this map
// has something to do with spaces. Give up.

// Check if it has exactly 1 space prefix or suffix.

// These spaces can be a padding,
// or a legitimate part of a key. Give up.

func (c *mapKeyChecker) warnWhitespace(key ast.Node) { _ = "STUB: not implemented"; return }

func (c *mapKeyChecker) warnDupKey(key ast.Node) { _ = "STUB: not implemented"; return }
