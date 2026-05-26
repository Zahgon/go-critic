package checkers

import (
	"go/ast"
	"regexp"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "whyNoLint"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Ensures that `//nolint` comments include an explanation"
	info.Before = `//nolint`
	info.After = `//nolint // reason`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForComment(&whyNoLintChecker{
			ctx: ctx,
			re:  regexp.MustCompile(`^// *nolint(?::[^ ]+)? *(.*)$`),
		}), nil
	})
}

type whyNoLintChecker struct {
	astwalk.WalkHandler

	ctx *linter.CheckerContext
	re  *regexp.Regexp
}

func (c whyNoLintChecker) VisitComment(cg *ast.CommentGroup) { _ = "STUB: not implemented"; return }
