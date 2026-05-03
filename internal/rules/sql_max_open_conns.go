package rules

import (
	"fmt"

	"github.com/IgorSteps/dblinter/internal/domain"
	"golang.org/x/tools/go/analysis"
)

type MaxOpenConnsRule struct {
	MaxOpenConnsRequired int
}

func (s *MaxOpenConnsRule) Check(pass *analysis.Pass, calls []domain.CallSite) error {
	// if call.Receiver.String() != "*sql.DB" || call.Method != "SetMaxOpenConns" {
	// 	return nil
	// }

	for _, call := range calls {
		fmt.Printf("receiver %s, method %s, args %v \n", call.Receiver.String(), call.Method, call.Args)
	}

	return nil
}

func NewMaxOpenConnsRuleFromConfig(config *domain.Config) *MaxOpenConnsRule {
	return &MaxOpenConnsRule{
		MaxOpenConnsRequired: config.MaxOpenConns,
	}
}
