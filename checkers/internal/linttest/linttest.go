package linttest

import (
	"go/ast"
	"go/token"
	"go/types"
	"runtime"
	"testing"

	"github.com/go-critic/go-critic/linter"

	"golang.org/x/tools/go/packages"
)

var sizes = types.SizesFor("gc", runtime.GOARCH)

func saneCheckersList(t *testing.T, checkers []*linter.CheckerInfo) []*linter.CheckerInfo {
	_ = "STUB: not implemented"
	return nil
}

// IntegrationTest specifies integration test options.
type IntegrationTest struct {
	Main string

	// Dir specifies a path to integration tests.
	Dir string
}

type CheckersTest struct {
	// IgnoreErrors is a checker names list those tests ignore parse/typecheck errors.
	IgnoreErrors []string
}

// Run executes every registered checker tests.
//
// TODO(quasilyte): document default options.
func (cfg *CheckersTest) Run(t *testing.T) { _ = "STUB: not implemented"; return }

// See #980.

type lintTarget struct {
	pattern      string
	ignoreErrors bool
}

func checkTarget(t *testing.T, target lintTarget, info *linter.CheckerInfo) {
	_ = "STUB: not implemented"
	return
}

func checkFile(t *testing.T, c *linter.Checker, ctx *linter.Context, f *ast.File) {
	_ = "STUB: not implemented"
	return
}

// stripDirectives replaces "///" comments with empty single-line
// comments, so the checkers that inspect comments see ordinary
// comment groups (with extra newlines, but that's not important).
func stripDirectives(f *ast.File) { _ = "STUB: not implemented"; return }

func getFilename(fset *token.FileSet, f *ast.File) string {
	_ = "STUB: not implemented"
	// see https://github.com/golang/go/issues/24498
	return ""
}

func checkUnmatched(ws warnings, matched map[*string]struct{}, t *testing.T, testFilename string) {
	_ = "STUB: not implemented"
	return
}

func newPackages(t *testing.T, pattern string, fset *token.FileSet) []*packages.Package {
	_ = "STUB: not implemented"
	return nil
}
