package checkers

import (
	"go/ast"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

const deprecatedPrefix = "Deprecated: "

func init() {
	var info linter.CheckerInfo
	info.Name = "deprecatedComment"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects malformed 'deprecated' doc-comments"
	info.Before = `
// deprecated, use FuncNew instead
func FuncOld() int`
	info.After = `
// Deprecated: use FuncNew instead
func FuncOld() int`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &deprecatedCommentChecker{ctx: ctx}

		c.commonPatterns = []string{
			"this type is deprecated",
			"this function is deprecated",
			"[[deprecated]]",
			"note: deprecated",
			"deprecated in",
			"deprecated. use",
			"deprecated! use",
			"deprecated use",
			// TODO(quasilyte): more of these?
		}

		// TODO(quasilyte): may want to generate this list programmatically.
		//
		// TODO(quasilyte): currently it only handles a single missing letter.
		// Might want to handle other kinds of common misspell/typo kinds.
		c.commonTypos = []string{
			"Dprecated: ",
			"Derecated: ",
			"Depecated: ",
			"Depekated: ",
			"Deprcated: ",
			"Depreated: ",
			"Deprected: ",
			"Deprecaed: ",
			"Deprecatd: ",
			"Deprecate: ",
			"Derpecate: ",
			"Derpecated: ",
			"Depreacted: ",
		}
		for i := range c.commonTypos {
			c.commonTypos[i] = strings.ToUpper(c.commonTypos[i])
		}

		return astwalk.WalkerForDocComment(c), nil
	})
}

type deprecatedCommentChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	commonPatterns []string
	commonTypos    []string
}

func (c *deprecatedCommentChecker) VisitDocComment(doc *ast.CommentGroup) {
	_ = "STUB: not implemented"
	// There are 3 accepted forms of deprecation comments:
	//
	// 1. inline, that can't be handled with a DocCommentVisitor.
	//    Note that "Deprecated: " may not even be the comment prefix there.
	//    Example: "The line number in the input. Deprecated: Kept for compatibility."
	//    TODO(quasilyte): fix it.
	//
	// 2. Longer form-1. It's a doc-comment that only contains "deprecation" notice.
	//
	// 3. Like form-2, but may also include doc-comment text.
	//    Distinguished by an empty line.
	//
	// See https://github.com/golang/go/issues/10909#issuecomment-136492606.
	//
	// It's desirable to see how people make mistakes with the format,
	// this is why there is currently no special treatment for these cases.
	// TODO(quasilyte): do more audits and grow the negative tests suite.
	//
	// TODO(quasilyte): there are also multi-line deprecation comments.
	return
}

// prev stores the previous line after it was trimmed.
// It's used to check whether the deprecation prefix is at the beginning of a new paragraph.

// TODO(quasilyte): handle multi-line doc comments.

// Check whether someone messed up with a prefix casing.

// Check is someone used comma instead of a colon.

// Check for other commonly used patterns.

// Detect some simple typos.

func (c *deprecatedCommentChecker) warnCasing(cause ast.Node, line string) {
	_ = "STUB: not implemented"
	return
}

func (c *deprecatedCommentChecker) warnPattern(cause ast.Node) { _ = "STUB: not implemented"; return }

func (c *deprecatedCommentChecker) warnComma(cause ast.Node) { _ = "STUB: not implemented"; return }

func (c *deprecatedCommentChecker) warnTypo(cause ast.Node, line string) {
	_ = "STUB: not implemented"
	return
}

func (c *deprecatedCommentChecker) warnParagraph(cause ast.Node) { _ = "STUB: not implemented"; return }
