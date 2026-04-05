package domain

import (
	"golang.org/x/tools/go/analysis"
)

type AnalyserRunner interface {
	Run(*analysis.Pass) (any, error)
}
