package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	dbLinter := Setup()
	// singlechecker.Main already handles OS exits internally.
	singlechecker.Main(dbLinter.Analyser)
}
