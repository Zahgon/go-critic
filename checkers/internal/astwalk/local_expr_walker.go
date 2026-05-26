package astwalk

import (
	"go/ast"
)

type localExprWalker struct {
	visitor LocalExprVisitor
}

func (w *localExprWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
