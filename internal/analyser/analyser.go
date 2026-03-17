package analyser

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var DbLinterAnalyzer = &analysis.Analyzer{
	Name: "dblinter",
	Doc:  "checks database against best practices",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	// fmt.Println("test")
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if stmt, ok := n.(*ast.AssignStmt); ok {
				fmt.Printf("stmt: %v\n", stmt)
				fmt.Printf("stmt.Rhs[0]: %v\n", stmt.Rhs[0])
				if callExpr, ok := stmt.Rhs[0].(*ast.CallExpr); ok {
					// fmt.Printf("callExpr: %v\n", callExpr)
					fun := callExpr.Fun
					if selectExpr, ok := fun.(*ast.SelectorExpr); ok {
						fmt.Printf("selectExpr: %v\n", selectExpr)
						var giveMeName *ast.Ident
						if tbd, ok := selectExpr.X.(*ast.Ident); ok {
							giveMeName = tbd
						}
						if giveMeName.Name == "sql" && selectExpr.Sel.Name == "Open" {
							pass.Reportf(stmt.Pos(), "found sql.Open")
						}
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
