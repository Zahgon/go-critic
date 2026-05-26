package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "dupImport"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects multiple imports of the same package under different aliases"
	info.Before = `
import (
	"fmt"
	printing "fmt" // Imported the second time
)`
	info.After = `
import(
	"fmt"
)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &dupImportChecker{ctx: ctx}, nil
	})
}

type dupImportChecker struct {
	ctx *linter.CheckerContext
}

func (c *dupImportChecker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

func (c *dupImportChecker) warn(importList []*ast.ImportSpec) { _ = "STUB: not implemented"; return }
