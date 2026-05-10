package main

import (
	"flag"

	"github.com/IgorSteps/dblinter/internal/adapters/koanfconfig"
	"github.com/IgorSteps/dblinter/internal/analysers"
	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/IgorSteps/dblinter/internal/rules"
	"golang.org/x/tools/go/analysis"
)

type DBLinter struct {
	Analyser *analysis.Analyzer
}

func Setup() (*DBLinter, error) {
	userConfigPath := ""
	flag.StringVar(&userConfigPath, "config", "", "optional path to config file")
	flag.Parse()
	config, err := koanfconfig.LoadConfig(userConfigPath)
	if err != nil {
		return nil, err
	}

	rules := []domain.Rule{rules.NewMaxOpenConnsRuleFromConfig(config.MaxOpenConns)}
	analyser := analysers.NewDBConnectionAnalyser(rules)

	return &DBLinter{
		Analyser: analyser,
	}, nil
}
