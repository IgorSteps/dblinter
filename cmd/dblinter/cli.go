package main

import (
	"github.com/IgorSteps/dblinter/internal/analysers"
	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/IgorSteps/dblinter/internal/rules"
	"golang.org/x/tools/go/analysis"
)

type DBLinter struct {
	Analyser *analysis.Analyzer
}

func Setup() *DBLinter {
	config := domain.NewConfig("10")
	rule := rules.NewMaxOpenConnsRuleFromConfig(config)
	analyser := analysers.NewDBConnectionAnalyser([]domain.Rule{rule})
	return &DBLinter{
		Analyser: analyser,
	}
}
