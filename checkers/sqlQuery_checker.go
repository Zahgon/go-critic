package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "sqlQuery"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects issue in Query() and Exec() calls"
	info.Before = `_, err := db.Query("UPDATE ...")`
	info.After = `_, err := db.Exec("UPDATE ...")`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&sqlQueryChecker{ctx: ctx}), nil
	})
}

type sqlQueryChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *sqlQueryChecker) VisitStmt(stmt ast.Stmt) { _ = "STUB: not implemented"; return }

// Query() has 2 return values.

// If Query() is called, but first return value is ignored,
// there is no way to close/read the returned rows.
// This can cause a connection leak.

func (c *sqlQueryChecker) funcIsQuery(funcExpr *ast.SelectorExpr) bool {
	_ = "STUB: not implemented"
	return false
}

// Stdlib and friends.

// sqlx.

// To avoid false positives (unrelated types can have Query method)
// check that the 1st returned type has Row-like name.

func (c *sqlQueryChecker) typeIsRowsLike(typ types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *sqlQueryChecker) funcIsExec(fn *types.Func) bool { _ = "STUB: not implemented"; return false }

// Expect exactly 2 results.

// Expect at least 1 param and it should be a string (query).

func (c *sqlQueryChecker) typeHasExecMethod(typ types.Type) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO(cristaloleg): is there something else to handle?

// Check embedded types.

func (c *sqlQueryChecker) warnAndSuggestExec(funcExpr *ast.SelectorExpr) {
	_ = "STUB: not implemented"
	return
}

func (c *sqlQueryChecker) warnRowsIgnored(funcExpr *ast.SelectorExpr) {
	_ = "STUB: not implemented"
	return
}
