package main

import (
	"github.com/IgorSteps/dblinter/internal/core"
	"github.com/IgorSteps/dblinter/internal/rules"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	setMaxOpenConnsRule := &core.Rule{
		Name:        "SetMaxOpenConns",
		Description: "SetMaxOpenConns set to required value",
	}

	setMaxOpenConnsRunner := &rules.SqlSetMaxOpenConnsAnaliserRunner{
		RequiredMaxOpenConns: "15",
	}

	setMaxOpenConnsAnalyser := core.NewAnalyser(setMaxOpenConnsRule, setMaxOpenConnsRunner)

	// singlechecker.Main already handles os exits internally.
	singlechecker.Main(setMaxOpenConnsAnalyser)
}
