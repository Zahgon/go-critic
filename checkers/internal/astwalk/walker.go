package astwalk

import (
	"go/types"

	"github.com/go-critic/go-critic/linter"
)

// WalkerForFuncDecl returns file walker implementation for FuncDeclVisitor.
func WalkerForFuncDecl(v FuncDeclVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForExpr returns file walker implementation for ExprVisitor.
func WalkerForExpr(v ExprVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForLocalExpr returns file walker implementation for LocalExprVisitor.
func WalkerForLocalExpr(v LocalExprVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForStmtList returns file walker implementation for StmtListVisitor.
func WalkerForStmtList(v StmtListVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForStmt returns file walker implementation for StmtVisitor.
func WalkerForStmt(v StmtVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForTypeExpr returns file walker implementation for TypeExprVisitor.
func WalkerForTypeExpr(v TypeExprVisitor, info *types.Info) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForLocalComment returns file walker implementation for LocalCommentVisitor.
func WalkerForLocalComment(v LocalCommentVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForComment returns file walker implementation for CommentVisitor.
func WalkerForComment(v CommentVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForDocComment returns file walker implementation for DocCommentVisitor.
func WalkerForDocComment(v DocCommentVisitor) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}

// WalkerForLocalDef returns file walker implementation for LocalDefVisitor.
func WalkerForLocalDef(v LocalDefVisitor, info *types.Info) linter.FileWalker {
	_ = "STUB: not implemented"
	return *new(linter.FileWalker)
}
