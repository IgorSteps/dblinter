package main

import (
	"flag"

	"github.com/IgorSteps/dblinter/internal/analysers"
	"github.com/IgorSteps/dblinter/internal/config"
	"github.com/IgorSteps/dblinter/internal/engine"
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
	config, err := config.LoadConfig(userConfigPath)
	if err != nil {
		return nil, err
	}

	maxOpenConsRule, err := rules.NewMaxOpenConnsRule(config.MaxOpenConns.Enabled, config.MaxOpenConns.Required)
	if err != nil {
		return nil, err
	}

	runner := engine.New([]rules.Rule{maxOpenConsRule})
	analyser := analysers.NewDBConnectionAnalyser(runner)

	return &DBLinter{
		Analyser: analyser,
	}, nil
}
