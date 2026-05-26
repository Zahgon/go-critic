package lintutil

import (
	"go/ast"
	"go/types"
)

// Different utilities to make simple analysis over typed ast values flow.
//
// It's primitive and can't replace SSA, but the bright side is that
// it does not require building an additional IR eagerly.
// Expected to be used sparingly inside a few checkers.
//
// If proven really useful, can be moved to go-toolsmith library.

// IsImmutable reports whether n can be modified through any operation.
func IsImmutable(info *types.Info, n ast.Expr) bool { _ = "STUB: not implemented"; return false }

// CouldBeMutated reports whether dst can be modified inside body.
//
// Note that it does not take already existing pointers to dst.
// An example of safe and correct usage is checking of something
// that was just defined, so the dst is a result of that definition.
func CouldBeMutated(info *types.Info, body ast.Node, dst ast.Expr) bool {
	_ = "STUB: not implemented"
	return false
	// Fast path.
}

// We don't track pass-by-value.
// If it's already a pointer, passing it by value
// means that there can be a potential indirect modification.
//
// It's possible to be less conservative here and find at least
// one such value pass before giving up.

// Identifier can be shadowed,
// so we need to check the object as well.

// Being conservative

// Address taken

// Incremented or decremented.
