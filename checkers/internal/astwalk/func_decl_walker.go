package astwalk

import (
	"go/ast"
)

type funcDeclWalker struct {
	visitor FuncDeclVisitor
}

func (w *funcDeclWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
