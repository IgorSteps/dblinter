package analysers

import (
	"github.com/IgorSteps/dblinter/internal/domain"
	"golang.org/x/tools/go/analysis"
)

func NewDBConnectionAnalyser(rules []domain.Rule) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "DBConnections",
		Doc:  "Checks database connections configurations follow best practices.",
		Run: func(pass *analysis.Pass) (any, error) {
			calls := domain.FindCallsSites(pass)

			for _, rule := range rules {
				err := rule.Check(pass, calls)
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		},
	}
}
