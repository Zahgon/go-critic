package astwalk

import (
	"go/ast"
)

type stmtWalker struct {
	visitor StmtVisitor
}

func (w *stmtWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
