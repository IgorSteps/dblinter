package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	dbLinter, err := Setup()
	if err != nil {
		fmt.Printf("failed to setup dblinter %v", err)
		os.Exit(1)
	}
	// singlechecker.Main already handles OS exits internally.
	singlechecker.Main(dbLinter.Analyser)
}
