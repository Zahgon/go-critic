package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "builtinShadowDecl"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects top-level declarations that shadow the predeclared identifiers"
	info.Before = `type int struct {}`
	info.After = `type myInt struct {}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &builtinShadowDeclChecker{ctx: ctx}, nil
	})
}

type builtinShadowDeclChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *builtinShadowDeclChecker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

// Don't check methods. They can shadow anything safely.

func (c *builtinShadowDeclChecker) visitGenDecl(decl *ast.GenDecl) {
	_ = "STUB: not implemented"
	return
}

func (c *builtinShadowDeclChecker) checkName(name *ast.Ident) { _ = "STUB: not implemented"; return }

func (c *builtinShadowDeclChecker) warn(ident *ast.Ident) { _ = "STUB: not implemented"; return }
