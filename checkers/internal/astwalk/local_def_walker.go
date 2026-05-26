package astwalk

import (
	"go/ast"
	"go/types"
)

type localDefWalker struct {
	visitor LocalDefVisitor
	info    *types.Info
}

func (w *localDefWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

func (w *localDefWalker) walkFunc(decl *ast.FuncDecl) { _ = "STUB: not implemented"; return }

func (w *localDefWalker) walkFuncBody(decl *ast.FuncDecl) { _ = "STUB: not implemented"; return }

// Multi-value assignment.
// Invariant: there is only 1 RHS.

// Simple 1-1 assignments.

// Decls always introduce new names.

// Ignore type/import specs

// var-specific decls without explicit init.

// var-specific decls that assign tuple results.

// Can be either var or const decl.

func (w *localDefWalker) walkSignature(decl *ast.FuncDecl) { _ = "STUB: not implemented"; return }
