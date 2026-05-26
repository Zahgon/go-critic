package linter

import (
	"regexp"
)

type checkerProto struct {
	info        *CheckerInfo
	constructor func(*Context) (*Checker, error)
}

// prototypes is a set of registered checkers that are not yet instantiated.
// Registration should be done with AddChecker function.
// Initialized checkers can be obtained with NewChecker function.
var prototypes = make(map[string]checkerProto)

func getCheckersInfo() []*CheckerInfo { _ = "STUB: not implemented"; return nil }

func addChecker(info *CheckerInfo, constructor func(*CheckerContext) (FileWalker, error)) {
	_ = "STUB: not implemented"
	return
}

// Validate param value type.

// OK.

func newChecker(ctx *Context, info *CheckerInfo) (*Checker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateCheckerInfo(info *CheckerInfo) error { _ = "STUB: not implemented"; return nil }

var validIdentRE = regexp.MustCompile(`^\w+$`)

func validateCheckerName(info *CheckerInfo) error { _ = "STUB: not implemented"; return nil }

func validateCheckerDocumentation(_ *CheckerInfo) error {
	_ = "STUB: not implemented"
	// TODO(quasilyte): validate documentation.
	return nil
}

func validateCheckerTags(info *CheckerInfo) error { _ = "STUB: not implemented"; return nil }
