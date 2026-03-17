package main

import (
	"github.com/IgorSteps/dblinter/internal/analyser"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyser.DbLinterAnalyzer)
}
