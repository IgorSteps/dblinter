package engine

import (
	"github.com/IgorSteps/dblinter/internal/diagnostics"
	"github.com/IgorSteps/dblinter/internal/rules"
)

type Runner struct {
	rules []rules.Rule
}

func New(rules []rules.Rule) *Runner {
	return &Runner{
		rules: rules,
	}
}

func (r *Runner) Run(calls []rules.CallSite) []diagnostics.Issue {
	var issues []diagnostics.Issue

	for _, rule := range r.rules {
		issues = append(issues, rule.Check(calls)...)
	}

	return issues
}
