package checkers

import (
	"go/ast"
	"regexp"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "regexpPattern"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious regexp patterns"
	info.Before = "regexp.MustCompile(`google.com|yandex.ru`)"
	info.After = "regexp.MustCompile(`google\\.com|yandex\\.ru`)"

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		domains := []string{
			"com",
			"org",
			"info",
			"net",
			"ru",
			"de",
		}

		allDomains := strings.Join(domains, "|")
		domainRE := regexp.MustCompile(`[^\\]\.(` + allDomains + `)\b`)
		return astwalk.WalkerForExpr(&regexpPatternChecker{
			ctx:      ctx,
			domainRE: domainRE,
		}), nil
	})
}

type regexpPatternChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	domainRE *regexp.Regexp
}

func (c *regexpPatternChecker) VisitExpr(x ast.Expr) { _ = "STUB: not implemented"; return }

func (c *regexpPatternChecker) warnDomain(cause ast.Expr, domain string) {
	_ = "STUB: not implemented"
	return
}
