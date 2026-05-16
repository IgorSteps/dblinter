package rules

import (
	"fmt"
	"go/ast"

	"github.com/IgorSteps/dblinter/internal/diagnostics"
)

const (
	requiredReceiver = "*database/sql.DB"
	requiredMethod   = "SetMaxOpenConns"
)

type MaxOpenConnsRule struct {
	Enabled  bool
	Required int
}

func (s *MaxOpenConnsRule) Check(calls []CallSite) []diagnostics.Issue {
	if !s.Enabled {
		return nil
	}
	var issues []diagnostics.Issue
	for _, call := range calls {
		if call.Receiver.String() == requiredReceiver && call.Method == requiredMethod {
			data, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				//return fmt.Errorf("argument to SetMaxOpenConns is not basic literal")
			}

			if data.Value != fmt.Sprint(s.Required) {
				issues = append(
					issues,
					diagnostics.Issue{
						RuleID:  "1",
						Pos:     call.Position,
						Message: fmt.Sprintf("MaxOpenConns: must be set to %d, but was set to %s", s.Required, data.Value),
					},
				)
			}
		}
	}

	return issues
}

func NewMaxOpenConnsRule(enabled bool, required int) (*MaxOpenConnsRule, error) {
	if required <= 0 {
		return nil, fmt.Errorf("required cannot be less than or equal to 0")
	}

	return &MaxOpenConnsRule{
		Enabled:  enabled,
		Required: required,
	}, nil
}
