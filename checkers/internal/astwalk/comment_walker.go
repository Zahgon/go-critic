package astwalk

import (
	"go/ast"
)

type commentWalker struct {
	visitor CommentVisitor
}

func (w *commentWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

func visitCommentGroups(cg *ast.CommentGroup, visit func(*ast.CommentGroup)) {
	_ = "STUB: not implemented"
	return
}
