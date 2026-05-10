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

	// TODO: this shouldn't happen here, move it elsewhere. But where?
	cfg, err := config.MaxOpenConns.ToDomain()
	if err != nil {
		return nil, err
	}

	rule := rules.NewMaxOpenConnsRuleFromConfig(cfg)
	analyser := analysers.NewDBConnectionAnalyser([]domain.Rule{rule})

	return &DBLinter{
		Analyser: analyser,
	}, nil
}
