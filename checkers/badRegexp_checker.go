package checkers

import (
	"go/ast"
	"unicode/utf8"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"github.com/quasilyte/regex/syntax"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "badRegexp"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious regexp patterns"
	info.Before = "regexp.MustCompile(`(?:^aa|bb|cc)foo[aba]`)"
	info.After = "regexp.MustCompile(`^(?:aa|bb|cc)foo[ab]`)"

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		opts := &syntax.ParserOptions{}
		c := &badRegexpChecker{
			ctx:    ctx,
			parser: syntax.NewParser(opts),
		}
		return astwalk.WalkerForExpr(c), nil
	})
}

type badRegexpChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	parser *syntax.Parser
	cause  ast.Expr

	flagStates  []regexpFlagState
	goodAnchors []syntax.Position
}

type regexpFlagState [utf8.RuneSelf]bool

func (c *badRegexpChecker) VisitExpr(x ast.Expr) { _ = "STUB: not implemented"; return }

func (c *badRegexpChecker) checkPattern(pat string) { _ = "STUB: not implemented"; return }

// In Go all flags (modifiers) are set to false by default,
// so we start from the empty flag set.

func (c *badRegexpChecker) markGoodCarets(e syntax.Expr) { _ = "STUB: not implemented"; return }

func (c *badRegexpChecker) walk(e syntax.Expr) { _ = "STUB: not implemented"; return }

// Creates a new context using the current context copy.
// New flags are evaluated inside a new context.
// After nested expressions are processed, previous context is restored.

// Like with OpGroupWithFlags, but doesn't evaluate any new flags.

func (c *badRegexpChecker) currentFlagState() *regexpFlagState {
	_ = "STUB: not implemented"
	return nil
}

func (c *badRegexpChecker) updateFlagState(state *regexpFlagState, e syntax.Expr, flagString string) {
	_ = "STUB: not implemented"
	return
}

// Should never happen in practice, but we don't want a panic

func (c *badRegexpChecker) checkNestedQuantifier(e syntax.Expr) { _ = "STUB: not implemented"; return }

func (c *badRegexpChecker) checkAltDups(alt syntax.Expr) {
	_ = "STUB: not implemented"
	// Seek duplicated alternation expressions.
	return
}

func (c *badRegexpChecker) isCharOrLit(e syntax.Expr) bool { _ = "STUB: not implemented"; return false }

func (c *badRegexpChecker) checkAltAnchor(alt syntax.Expr) {
	_ = "STUB: not implemented"
	// Seek suspicious anchors.
	return
}

// Case 1: an alternation of literals where 1st expr begins with ^ anchor.

// Case 2: an alternation of literals where last expr ends with $ anchor.

func (c *badRegexpChecker) checkCharClassRanges(cc syntax.Expr) bool {
	_ = "STUB: not implemented"
	// Seek for suspicious ranges like `!-_`.
	//
	// We permit numerical ranges (0-9, hex and octal literals)
	// and simple ascii letter ranges.
	return false
}

func (c *badRegexpChecker) checkCharClassDups(cc syntax.Expr) {
	_ = "STUB: not implemented"
	// Seek for excessive elements inside a character class.
	// Report them as intersections.
	return
}

// Can't had duplicates.

// 1. Collect ranges, O(n).

// How to cover all symbols?

// 9-10
// 12-13
// 32

// 48-57
// 65-90
// 95
// 97-122

// Give up: unknown escape sequence.

// Give up: unexpected operation inside char class.

// 2. Sort ranges, O(nlogn).

// 3. Search for duplicates, O(n).

func (c *badRegexpChecker) charClassBoundRune(e syntax.Expr) rune {
	_ = "STUB: not implemented"
	return 0
}

func (c *badRegexpChecker) octalToRune(e syntax.Expr) rune { _ = "STUB: not implemented"; return 0 }

func (c *badRegexpChecker) hexToRune(e syntax.Expr) rune { _ = "STUB: not implemented"; return 0 }

func (c *badRegexpChecker) stringToRune(s string) rune { _ = "STUB: not implemented"; return 0 }

func (c *badRegexpChecker) addGoodAnchor(pos syntax.Position) { _ = "STUB: not implemented"; return }

func (c *badRegexpChecker) isGoodAnchor(e syntax.Expr) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *badRegexpChecker) warnf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *badRegexpChecker) warnSloppyCharRange(rng, charClass string) {
	_ = "STUB: not implemented"
	return
}

func (c *badRegexpChecker) warnCharClassDup(x, y, charClass string) {
	_ = "STUB: not implemented"
	return
}
