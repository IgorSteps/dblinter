package analysers

import (
	"github.com/IgorSteps/dblinter/internal/engine"
	"golang.org/x/tools/go/analysis"
)

func NewDBConnectionAnalyser(runner *engine.Runner) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "DBConnections",
		Doc:  "Checks database connections configurations follow best practices.",
		Run: func(pass *analysis.Pass) (any, error) {
			calls := FindCallsSites(pass)

			issues := runner.Run(calls)

			for _, issue := range issues {
				pass.Reportf(issue.Pos, "%s", issue.Message)
			}

			return nil, nil
		},
	}
}
