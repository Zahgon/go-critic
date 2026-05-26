package astwalk

import (
	"go/ast"
	"go/types"
)

type typeExprWalker struct {
	visitor TypeExprVisitor
	info    *types.Info
}

func (w *typeExprWalker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

func (w *typeExprWalker) visit(x ast.Expr) bool { _ = "STUB: not implemented"; return false }

func (w *typeExprWalker) walk(x ast.Node) bool { _ = "STUB: not implemented"; return false }

// Pointer conversions require parenthesis around pointer type.
// These casts are represented as call expressions.
// Because it's impossible for the visitor to distinguish such
// "required" parenthesis, walker skips outmost parenthesis in such cases.

// Like with conversions, method expressions are another special.

// Embedded interface.

func (w *typeExprWalker) inspectInner(x ast.Expr) bool { _ = "STUB: not implemented"; return false }

func (w *typeExprWalker) walkSignature(typ *ast.FuncType) { _ = "STUB: not implemented"; return }
