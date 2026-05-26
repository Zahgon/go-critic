package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/linter"

	"github.com/quasilyte/go-ruleguard/ruleguard"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "ruleguard"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"rules": {
			Value: "",
			Usage: "comma-separated list of gorule file paths. Glob patterns such as 'rules-*.go' may be specified",
		},
		"debug": {
			Value: "",
			Usage: "enable debug for the specified named rules group",
		},
		"failOnError": {
			Value: false,
			Usage: "deprecated, use failOn param; if set to true, identical to failOn='all', otherwise failOn=''",
		},
		"failOn": {
			Value: "",
			Usage: `Determines the behavior when an error occurs while parsing ruleguard files.
If flag is not set, log error and skip rule files that contain an error.
If flag is set, the value must be a comma-separated list of error conditions.
* 'import': rule refers to a package that cannot be loaded.
* 'dsl':    gorule file does not comply with the ruleguard DSL.`,
		},
		"enable": {
			Value: "<all>",
			Usage: "comma-separated list of enabled groups or skip empty to enable everything",
		},
		"disable": {
			Value: "",
			Usage: "comma-separated list of disabled groups or skip empty to enable everything",
		},
	}
	info.Summary = "Runs user-defined rules using ruleguard linter"
	info.Details = "Reads a rules file and turns them into go-critic checkers."
	info.Before = `N/A`
	info.After = `N/A`
	info.Note = "See https://github.com/quasilyte/go-ruleguard."

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return newRuleguardChecker(&info, ctx)
	})
}

// parseErrorHandler is used to determine whether to ignore or fail ruleguard parsing errors.
type parseErrorHandler struct {
	// failureConditions is a map of predicates which are evaluated against a ruleguard parsing error.
	// If at least one predicate returns true, then an error is returned.
	// Otherwise, the ruleguard file is skipped.
	failureConditions map[string]func(err error) bool
}

// failOnParseError returns true if a parseError occurred and that error should be not be ignored.
func (e parseErrorHandler) failOnParseError(parseError error) bool {
	_ = "STUB: not implemented"
	return false
}

func newErrorHandler(failOnErrorFlag string) (*parseErrorHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wrong flag value.

func newRuleguardChecker(info *linter.CheckerInfo, ctx *linter.CheckerContext) (*ruleguardChecker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The only possible returned error is ErrBadPattern, when pattern is malformed.

type ruleguardChecker struct {
	ctx *linter.CheckerContext

	debugGroup string
	engine     *ruleguard.Engine
}

func (c *ruleguardChecker) WalkFile(f *ast.File) { _ = "STUB: not implemented"; return }

func runRuleguardEngine(ctx *linter.CheckerContext, f *ast.File, e *ruleguard.Engine, runCtx *ruleguard.RunContext) {
	_ = "STUB: not implemented"
	return
}

// TODO(quasilyte): investigate whether we should add a rule name as
// a message prefix here.

// Normally this should never happen, but since
// we don't have a better mechanism to report errors,
// emit a warning.

func debugPrint(s string) { _ = "STUB: not implemented"; return }
