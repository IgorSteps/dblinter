package main

import (
	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/IgorSteps/dblinter/internal/rules"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	ruleInfo := &domain.RuleInfo{
		Name:        "SetMaxOpenConns",
		Description: "SetMaxOpenConns set to required value",
	}

	rule := rules.NewMaxOpenConnsRule(ruleInfo, &domain.Analyser{})
	analyser := domain.NewAnalyser(ruleInfo, rule)
	// singlechecker.Main already handles OS exits internally.
	singlechecker.Main(analyser)
}
