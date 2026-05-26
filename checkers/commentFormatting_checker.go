package checkers

import (
	"go/ast"
	"regexp"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "commentFormatting"
	info.Tags = []string{linter.StyleTag}
	info.Summary = "Detects comments with non-idiomatic formatting"
	info.Before = `//This is a comment`
	info.After = `// This is a comment`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		regexpPatterns := []*regexp.Regexp{
			regexp.MustCompile(`^//[\w-]+:.*$`), // e.g.: key: value
		}
		equalPatterns := []string{
			"//nolint",
		}
		parts := []string{
			"//go:generate ",   // e.g.: go:generate value
			"//line /",         // e.g.: line /path/to/file:123
			"//nolint ",        // e.g.: nolint
			"//noinspection ",  // e.g.: noinspection ALL, some GoLand and friends versions
			"//region",         // e.g.: region awawa, used by GoLand and friends for custom folding
			"//endregion",      // e.g.: endregion awawa or endregion, closes GoLand regions
			"//<editor-fold",   // e.g.: <editor-fold desc="awawa"> or <editor-fold>, used by VSCode for custom folding
			"//</editor-fold>", // e.g.: </editor-fold>, closes VSCode regions
			"//export ",        // e.g.: export Foo
			"///",              // e.g.: vertical breaker /////////////
			"//+",
			"//#",
			"//-",
			"//!",
		}

		return astwalk.WalkerForComment(&commentFormattingChecker{
			ctx:            ctx,
			partPatterns:   parts,
			equalPatterns:  equalPatterns,
			regexpPatterns: regexpPatterns,
		}), nil
	})
}

type commentFormattingChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	partPatterns   []string
	equalPatterns  []string
	regexpPatterns []*regexp.Regexp
}

func (c *commentFormattingChecker) VisitComment(cg *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return
}

// Make a decision based on a first comment text rune.

func (c *commentFormattingChecker) specialChar(r rune) bool {
	_ = "STUB: not implemented"
	// Permitted list to avoid false-positives.
	return false
}

func (c *commentFormattingChecker) warn(comment *ast.Comment) { _ = "STUB: not implemented"; return }
