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
	Enabled  bool
	Required int
}

func (s *MaxOpenConnsRule) Check(pass *analysis.Pass, calls []domain.CallSite) error {
	for _, call := range calls {
		if call.Receiver.String() == requiredReceiver && call.Method == requiredMethod {
			data, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				return fmt.Errorf("argument to SetMaxOpenConns is not basic literal")
			}

			if data.Value != fmt.Sprint(s.Required) {
				pass.Reportf(call.Position, "MaxOpenConns: must be set to %d, but was set to %s", s.Required, data.Value)
			}
		}
	}

	return nil
}

func NewMaxOpenConnsRuleFromConfig(config *domain.MaxOpenConnsConfig) *MaxOpenConnsRule {
	return &MaxOpenConnsRule{
		Enabled:  config.Enabled,
		Required: config.Required,
	}
}
