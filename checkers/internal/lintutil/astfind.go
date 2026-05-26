package lintutil

import (
	"go/ast"
)

// FindNode applies pred for root and all it's childs until it returns true.
// If followFunc is defined, it's called before following any node to check whether it needs to be followed.
// followFunc has to return true in order to continuing traversing the node and return false otherwise.
// Matched node is returned.
// If none of the nodes matched predicate, nil is returned.
func FindNode(root ast.Node, followFunc, pred func(ast.Node) bool) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

// ContainsNode reports whether `FindNode(root, pred)!=nil`.
func ContainsNode(root ast.Node, pred func(ast.Node) bool) bool {
	_ = "STUB: not implemented"
	return false
}
