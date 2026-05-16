package analysers

import (
	"errors"

	"github.com/IgorSteps/dblinter/internal/engine"
	"golang.org/x/tools/go/analysis"
)

func NewDBConnectionAnalyser(runner *engine.Runner) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "DBConnections",
		Doc:  "Checks database connections configurations follow best practices.",
		Run: func(pass *analysis.Pass) (any, error) {
			calls := FindCallsSites(pass)

			issues, errs := runner.Run(calls)
			if errs != nil {
				return nil, errors.Join(errs...)
			}

			for _, issue := range issues {
				pass.Reportf(issue.Pos, "msg: %s, doc: %s", issue.Message, issue.Doc)
			}

			return nil, nil
		},
	}
}
