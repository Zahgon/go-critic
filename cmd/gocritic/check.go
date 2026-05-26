package main

import (
	"bytes"
	"context"
	"flag"
	"go/ast"
	"go/token"
	"regexp"

	"github.com/go-critic/go-critic/linter"

	"golang.org/x/tools/go/packages"
)

// Main implements sub-command entry point.
func runCheck(_ context.Context, args []string) error { _ = "STUB: not implemented"; return nil }

type program struct {
	ctx *linter.Context

	flagSet *flag.FlagSet

	fset *token.FileSet

	loadedPackages []*packages.Package

	infoList []*linter.CheckerInfo

	checkers []*linter.Checker

	packages []string

	foundIssues bool

	checkerParams boundCheckerParams

	filters struct {
		enableAll       bool
		enable          []string
		disable         []string
		defaultCheckers []string
	}

	workDir string
	gopath  string
	goroot  string

	cpuProfile string
	memProfile string

	cpuProfileData bytes.Buffer

	goVersion          string
	concurrency        int
	exitCode           int
	checkTests         bool
	checkGenerated     bool
	shorterErrLocation bool
	verbose            bool
}

func (p *program) exit() error { _ = "STUB: not implemented"; return nil }

func (p *program) runCheckers() error { _ = "STUB: not implemented"; return nil }

func (p *program) checkPackage(pkg *packages.Package) { _ = "STUB: not implemented"; return }

func (p *program) checkFile(f *ast.File) { _ = "STUB: not implemented"; return }

// All checkers are expected to use *lint.Context
// as read-only structure, so no copying is required.

// Checker signals unexpected error with panic(error).

// There were no panic

// Some other kind of run-time panic.
// Undo the recover and resume panic.

func (p *program) initCheckers() error { _ = "STUB: not implemented"; return nil }

func (p *program) loadProgram() error { _ = "STUB: not implemented"; return nil }

type boundCheckerParams struct {
	ints    map[string]*int
	bools   map[string]*bool
	strings map[string]*string
}

// bindCheckerParams registers command-line flags for every checker parameter.
func (p *program) bindCheckerParams() error { _ = "STUB: not implemented"; return nil }

// Checked in AddChecker

func (p *program) checkerParamKey(info *linter.CheckerInfo, pname string) string {
	_ = "STUB: not implemented"
	return ""
}

// bindDefaultEnabledList calculates the default value for -enable param.
func (p *program) bindDefaultEnabledList() error { _ = "STUB: not implemented"; return nil }

func (p *program) parseArgs(args []string) error { _ = "STUB: not implemented"; return nil }

func addTrailingSlash(s string) string { _ = "STUB: not implemented"; return "" }

func (p *program) startProfiling() error { _ = "STUB: not implemented"; return nil }

func (p *program) finishProfiling() error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // 0o666 is okay

// get up-to-date statistics

// assignCheckerParams initializes checker parameter values using
// values that are coming from the command-line arguments.
func (p *program) assignCheckerParams() error { _ = "STUB: not implemented"; return nil }

// Checked in AddChecker

var generatedFileCommentRE = regexp.MustCompile("Code generated .* DO NOT EDIT.")

func (p *program) isGenerated(f *ast.File) bool { _ = "STUB: not implemented"; return false }

func (p *program) getFilename(f *ast.File) string {
	_ = "STUB: not implemented"
	// See https://github.com/golang/go/issues/24498.
	return ""
}

func (p *program) shortenLocation(loc string) string {
	_ = "STUB: not implemented"
	// If possible, construct relative path.
	return ""
}

// Return the representation that is shorter.
