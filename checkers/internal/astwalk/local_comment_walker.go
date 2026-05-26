package astwalk

import (
	"go/ast"
)

type localCommentWalker struct {
	visitor LocalCommentVisitor
}

func (w *localCommentWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

// Not sure that decls/comments are sorted
// by positions, so do a naive full scan for now.
