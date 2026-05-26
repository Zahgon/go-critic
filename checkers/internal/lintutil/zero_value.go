package lintutil

import (
	"go/ast"
	"go/types"
)

// IsZeroValue reports whether x represents zero value of its type.
//
// The functions is conservative and may return false for zero values
// if some cases are not handled in a comprehensive way
// but is should never return true for something that's not a proper zv.
func IsZeroValue(info *types.Info, x ast.Expr) bool { _ = "STUB: not implemented"; return false }

// Note that this function is not comprehensive.

// ZeroValueOf returns a zero value expression for typeExpr of type typ.
// If function can't find such a value, nil is returned.
func ZeroValueOf(typeExpr ast.Expr, typ types.Type) ast.Expr {
	_ = "STUB: not implemented"
	return *new(ast.Expr)
}

func isDefaultLiteralType(typ types.Type) bool { _ = "STUB: not implemented"; return false }
