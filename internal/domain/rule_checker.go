package domain

import "golang.org/x/tools/go/analysis"

type Rule interface {
	Check(pass *analysis.Pass, calls []CallSite) error
}
