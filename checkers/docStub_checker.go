package checkers

import (
	"go/ast"
	"regexp"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "docStub"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects comments that silence go lint complaints about doc-comment"
	info.Before = `
// Foo ...
func Foo() {
}`
	info.After = `
// (A) - remove the doc-comment stub
func Foo() {}
// (B) - replace it with meaningful comment
// Foo is a demonstration-only function.
func Foo() {}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		re := `(?i)^\.\.\.$|^\.$|^xxx\.?$|^whatever\.?$`
		c := &docStubChecker{
			ctx:           ctx,
			stubCommentRE: regexp.MustCompile(re),
		}
		return c, nil
	})
}

type docStubChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	stubCommentRE *regexp.Regexp
}

func (c *docStubChecker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

// Only 1 spec, use doc from the decl itself.

// N specs, use per-spec doc.

func (c *docStubChecker) visitDoc(decl ast.Node, sym *ast.Ident, doc *ast.CommentGroup, article bool) {
	_ = "STUB: not implemented"
	return
}

// Skip optional article.

// Now try to detect the "stub" part.

func (c *docStubChecker) warn(cause ast.Node) { _ = "STUB: not implemented"; return }
