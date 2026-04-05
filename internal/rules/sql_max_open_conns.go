package rules

import (
	"fmt"
	"log"

	"github.com/IgorSteps/dblinter/internal/domain"
	"golang.org/x/tools/go/analysis"
)

var (
	functionList = []*domain.FunctionCall{
		{
			Package: "sql",
			Name:    "SetMaxOpenConns",
		},
	}
)

type MaxOpenConnsRule struct {
	RuleInfo *domain.RuleInfo
	Analyser *domain.Analyser
}

func (s *MaxOpenConnsRule) Run(pass *analysis.Pass) (any, error) {
	targets, err := s.Analyser.Run(pass, functionList)
	if err != nil {
		log.Printf("max open conns rule failed: %v", err.Error())
		return nil, err
	}

	validationFailuresCounter := 0
	for _, target := range targets {
		if target.Arg != "15" {
			validationFailuresCounter++
		}
	}

	if validationFailuresCounter != 0 {
		return nil, fmt.Errorf("rule not met")
	}
	return nil, nil
}

func NewMaxOpenConnsRule(ruleInfo *domain.RuleInfo, analyser *domain.Analyser) *MaxOpenConnsRule {
	return &MaxOpenConnsRule{
		RuleInfo: ruleInfo,
		Analyser: analyser,
	}
}
