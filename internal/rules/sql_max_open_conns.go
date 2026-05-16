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
	ID       string
	Doc      string
	Enabled  bool
	Required int
}

func (s *MaxOpenConnsRule) Check(calls []CallSite) ([]diagnostics.Issue, error) {
	if !s.Enabled {
		return nil, nil
	}
	var issues []diagnostics.Issue
	for _, call := range calls {
		if call.Receiver.String() == requiredReceiver && call.Method == requiredMethod {
			// TODO: we shouldn't be doing ast operations inside rules, remove.
			data, ok := call.Args[0].(*ast.BasicLit)
			if !ok {
				// TODO: remove this error case possibility.
				return []diagnostics.Issue{}, fmt.Errorf("argument to SetMaxOpenConns is not basic literal")
			}

			if data.Value != fmt.Sprint(s.Required) {
				issues = append(
					issues,
					diagnostics.Issue{
						RuleID:  s.ID,
						Doc:     s.Doc,
						Pos:     call.Position,
						Message: fmt.Sprintf("MaxOpenConns: must be set to %d, but was set to %s", s.Required, data.Value),
					},
				)
			}
		}
	}

	return issues, nil
}

func NewMaxOpenConnsRule(enabled bool, required int) (*MaxOpenConnsRule, error) {
	if required <= 0 {
		return nil, fmt.Errorf("required cannot be less than or equal to 0")
	}

	return &MaxOpenConnsRule{
		ID:       "DBL-1",
		Doc:      "TBD",
		Enabled:  enabled,
		Required: required,
	}, nil
}
