package rules

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

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
		if call.Receiver.String() != requiredReceiver {
			continue
		}

		if call.Method != requiredMethod {
			continue
		}

		if len(call.Args) == 0 {
			continue
		}

		arg, ok := call.Args[0].(*ast.BasicLit)
		if !ok {
			return []diagnostics.Issue{}, fmt.Errorf("expected argument to be basic literal, but got %v", arg)
		}

		if arg.Kind != token.INT {
			return []diagnostics.Issue{}, fmt.Errorf("expected argument to be int, but got %s", arg.Kind.String())
		}

		actualValue, err := strconv.Atoi(arg.Value)
		if err != nil {
			return []diagnostics.Issue{}, err
		}

		if actualValue != s.Required {
			issues = append(
				issues,
				diagnostics.Issue{
					RuleID:  s.ID,
					Doc:     s.Doc,
					Pos:     call.Position,
					Message: fmt.Sprintf("MaxOpenConns: must be set to %d, but was set to %d", s.Required, actualValue),
				},
			)
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
