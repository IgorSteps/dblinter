package main

import (
	"github.com/IgorSteps/dblinter/internal/analysers"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	// singlechecker.Main already handles OS exits internally.
	singlechecker.Main(analysers.MyMaxOpenConnsAnalasyer)
}
