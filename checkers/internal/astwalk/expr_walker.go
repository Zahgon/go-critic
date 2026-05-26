package astwalk

import (
	"go/ast"
)

type exprWalker struct {
	visitor ExprVisitor
}

func (w *exprWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
