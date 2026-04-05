package domain

import (
	"golang.org/x/tools/go/analysis"
)

func NewAnalyser(rule *RuleInfo, runner AnalyserRunner) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: rule.Name,
		Doc:  rule.Description,
		Run:  runner.Run,
	}
}
