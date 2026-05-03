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

func Setup() (*DBLinter, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}
	rule := rules.NewMaxOpenConnsRuleFromConfig(&config.MaxOpenConnectionsRuleConfig)
	analyser := analysers.NewDBConnectionAnalyser([]domain.Rule{rule})
	return &DBLinter{
		Analyser: analyser,
	}, nil
}
