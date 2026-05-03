package rules

import (
	"fmt"
	"go/ast"

	"github.com/IgorSteps/dblinter/internal/domain"
	"golang.org/x/tools/go/analysis"
)

const (
	requiredReceiver = "*database/sql.DB"
	requiredMethod   = "SetMaxOpenConns"
)

type MaxOpenConnsRule struct {
	MaxOpenConnsRequired string
}

func (s *MaxOpenConnsRule) Check(pass *analysis.Pass, calls []domain.CallSite) error {
	for _, call := range calls {
		if call.Receiver.String() == requiredReceiver && call.Method == requiredMethod {
			data, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				return fmt.Errorf("argument to SetMaxOpenConns is not basic literal")
			}

			if data.Value != s.MaxOpenConnsRequired {
				pass.Reportf(call.Position, "MaxOpenConns must be set to %s, but was set to %s", s.MaxOpenConnsRequired, data.Value)
			}
		}
	}

	return nil
}

func NewMaxOpenConnsRuleFromConfig(config *domain.Config) *MaxOpenConnsRule {
	return &MaxOpenConnsRule{
		MaxOpenConnsRequired: config.MaxOpenConns,
	}
}
