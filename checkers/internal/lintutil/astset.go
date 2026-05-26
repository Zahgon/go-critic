package lintutil

import (
	"go/ast"
)

// AstSet is a simple ast.Node set.
// Zero value is ready to use set.
// Can be reused after Clear call.
type AstSet struct {
	items []ast.Node
}

// Contains reports whether s contains x.
func (s *AstSet) Contains(x ast.Node) bool { _ = "STUB: not implemented"; return false }

// Insert pushes x in s if it's not already there.
// Returns true if element was inserted.
func (s *AstSet) Insert(x ast.Node) bool { _ = "STUB: not implemented"; return false }

// Clear removes all element from set.
func (s *AstSet) Clear() { _ = "STUB: not implemented"; return }

// Len returns the number of elements contained inside s.
func (s *AstSet) Len() int { _ = "STUB: not implemented"; return 0 }
