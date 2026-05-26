package analyzer

import (
	"sync"

	_ "github.com/go-critic/go-critic/checkers" // Register go-critic checkers
	"github.com/go-critic/go-critic/linter"

	"golang.org/x/tools/go/analysis"
)

type gocritic struct {
	infoList  []*linter.CheckerInfo
	goVersion linter.GoVersion
}

var (
	globalGocriticMu        sync.Mutex
	globalGocritic          *gocritic
	globalInitErrorReported bool
)

func runAnalyzer(pass *analysis.Pass) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func asDiag(c *linter.Checker, warning linter.Warning) analysis.Diagnostic {
	_ = "STUB: not implemented"
	return *new(analysis.Diagnostic)
}

// prepareGocritic initializes a new gocritic object,
// but unlike newGocritic() it could use a cached version.
func prepareGocritic() (*gocritic, error) { _ = "STUB: not implemented"; return nil, nil }

// Don't report init error ever again if it was already reported.

func newGocritic() (*gocritic, error) { _ = "STUB: not implemented"; return nil, nil }

// Checked in AddChecker

func filterCheckersList(infoList []*linter.CheckerInfo) []*linter.CheckerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (critic *gocritic) createCheckers(ctx *linter.Context) ([]*linter.Checker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
