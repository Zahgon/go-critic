package linttest

import (
	"io"
	"regexp"
)

var (
	warningDirectiveRE = regexp.MustCompile(`^\s*/\*! (.*) \*/`)
)

type warnings map[int][]string

func newWarnings(r io.Reader) (warnings, error) {
	_ = "STUB: not implemented"
	return *new(warnings), nil
}

func (ws warnings) find(line int, text string) *string { _ = "STUB: not implemented"; return nil }
