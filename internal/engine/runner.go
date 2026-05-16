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

func (r *Runner) Run(calls []rules.CallSite) ([]diagnostics.Issue, []error) {
	var issues []diagnostics.Issue
	var errors []error

	for _, rule := range r.rules {
		issues, err := rule.Check(calls)
		if err != nil {
			// Prefer to collate all the errors rather than exiting early,
			// so that the user doesn't stuck in a fix/rerun cycle.
			errors = append(errors, err)
			continue
		}
		issues = append(issues, issues...)
	}

	return issues, nil
}
