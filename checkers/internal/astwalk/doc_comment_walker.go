package astwalk

import (
	"go/ast"
)

type docCommentWalker struct {
	visitor DocCommentVisitor
}

func (w *docCommentWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
