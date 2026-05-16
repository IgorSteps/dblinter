package rules

import (
	"github.com/IgorSteps/dblinter/internal/diagnostics"
)

type Rule interface {
	Check(calls []CallSite) ([]diagnostics.Issue, error)
}
