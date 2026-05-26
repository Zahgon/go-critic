package checkers

import (
	"go/ast"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"github.com/quasilyte/regex/syntax"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "regexpSimplify"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag, linter.OpinionatedTag}
	info.Summary = "Detects regexp patterns that can be simplified"
	info.Before = "regexp.MustCompile(`(?:a|b|c)   [a-z][a-z]*`)"
	info.After = "regexp.MustCompile(`[abc] {3}[a-z]+`)"

	// TODO(quasilyte): add params to control most opinionated replacements
	// like `[0-9] -> \d`
	//      `[[:digit:]] -> \d`
	//      `[A-Za-z0-9_]` -> `\w`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		opts := &syntax.ParserOptions{
			NoLiterals: true,
		}
		c := &regexpSimplifyChecker{
			ctx:    ctx,
			parser: syntax.NewParser(opts),
			out:    &strings.Builder{},
		}
		return astwalk.WalkerForExpr(c), nil
	})
}

type regexpSimplifyChecker struct {
	astwalk.WalkHandler
	ctx    *linter.CheckerContext
	parser *syntax.Parser

	// out is a tmp buffer where we build a simplified regexp pattern.
	out *strings.Builder
	// score is a number of applied simplifications
	score int
}

func (c *regexpSimplifyChecker) VisitExpr(x ast.Expr) { _ = "STUB: not implemented"; return }

// Skip scary regexp patterns for now.

// Only do 2 passes.

func (c *regexpSimplifyChecker) simplify(pass int, pat string) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO(quasilyte): suggest char ranges for things like [012345689]?
// TODO(quasilyte): evaluate char range to suggest better replacements.
// TODO(quasilyte): (?:ab|ac) -> a[bc]
// TODO(quasilyte): suggest "s" and "." flag if things like [\w\W] are used.
// TODO(quasilyte): x{n}x? -> x{n,n+1}

// This happens only in one of two cases:
// 1. Parser has a bug and we got invalid AST for the given pattern.
// 2. Simplifier incorrectly built a replacement string from the AST.

func (c *regexpSimplifyChecker) walk(e syntax.Expr) { _ = "STUB: not implemented"; return }

// TODO(quasilyte): is it worth it to analyze repeat argument
// more closely and handle `{n,n} -> {n}` cases?

// Maybe {0} should be reported by another check, regexpLint?

func (c *regexpSimplifyChecker) walkGroup(g syntax.Expr) { _ = "STUB: not implemented"; return }

func (c *regexpSimplifyChecker) simplifyNegCharClass(e syntax.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *regexpSimplifyChecker) simplifyCharClass(e syntax.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

// Can't take outside of the char group without escaping.

func (c *regexpSimplifyChecker) canMerge(x, y syntax.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *regexpSimplifyChecker) canCombine(x, y syntax.Expr) (threshold int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (c *regexpSimplifyChecker) concatLiteral(e syntax.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *regexpSimplifyChecker) allChars(e syntax.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *regexpSimplifyChecker) factorPrefixSuffix(alt syntax.Expr) bool {
	_ = "STUB: not implemented"
	// TODO: more forms of prefixes/suffixes?
	//
	// A more generalized algorithm could handle `fo|fo1|fo2` -> `fo[12]?`.
	// but it's an open question whether the latter form universally better.
	//
	// Right now it handles only the simplest cases:
	// `http|https` -> `https?`
	// `xfoo|foo` -> `x?foo`
	return false
}

// Reject non-literals and identical strings early

// Let x be a shorter string.

// Do we have a common prefix?

// Do we have a common suffix?

func (c *regexpSimplifyChecker) walkAlt(alt syntax.Expr) {
	_ = "STUB: not implemented"
	// `x|y|z` -> `[xyz]`.
	return
}

func (c *regexpSimplifyChecker) walkConcat(concat syntax.Expr) { _ = "STUB: not implemented"; return }

// Try merging `xy*` into `x+` where x=y.

// Try combining `xy` into `x{2}` where x=y.

// Can combine at least 1 pair.

func (c *regexpSimplifyChecker) simplifyCharRange(rng syntax.Expr) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *regexpSimplifyChecker) warn(cause ast.Expr, orig, suggest string) {
	_ = "STUB: not implemented"
	return
}
