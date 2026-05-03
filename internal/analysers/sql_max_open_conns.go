package analysers

import (
	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/IgorSteps/dblinter/internal/rules"
	"golang.org/x/tools/go/analysis"
)

var MyMaxOpenConnsAnalasyer = &analysis.Analyzer{
	Name: "MyMaxOpenConnsAnalyser",
	Doc:  "Finds the thing",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	calls := domain.FindCallsSites(pass)

	myRule := rules.MaxOpenConnsRule{
		MaxOpenConnsRequired: 10,
	}

	return "", myRule.Check(pass, calls)
}
