package checkers

import (
	"go/ast"
	"regexp"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "commentedOutImport"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects commented-out imports"
	info.Before = `
import (
	"fmt"
	//"os"
)`
	info.After = `
import (
	"fmt"
)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		const pattern = `(?m)^(?://|/\*)?\s*"([a-zA-Z0-9_/]+)"\s*(?:\*/)?$`
		return &commentedOutImportChecker{
			ctx:            ctx,
			importStringRE: regexp.MustCompile(pattern),
		}, nil
	})
}

type commentedOutImportChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	importStringRE *regexp.Regexp
}

func (c *commentedOutImportChecker) WalkFile(f *ast.File) {
	_ = "STUB: not implemented"
	// TODO(quasilyte): handle commented-out import spec,
	// for example: // import "errors".
	return
}

// Import decls can only be in the beginning of the file.
// If we've met some other decl, there will be no more
// import decls.

// Find comments inside this import decl span.

// Below the decl, stop.

// Before the decl, skip.

func (c *commentedOutImportChecker) warn(cause ast.Node, path string) {
	_ = "STUB: not implemented"
	return
}
