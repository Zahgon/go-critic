package astwalk

import (
	"go/ast"
)

type stmtListWalker struct {
	visitor StmtListVisitor
}

func (w *stmtListWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }
