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
	requiredVal := "15"
	found := false
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// db.SetMaxOpenWhatever is a ExprStmt
			// as such confirm node is indeed an ExprStmt
			exprNode, ok := node.(*ast.ExprStmt)
			if !ok {
				return true
			}
			// assert that the node is a function call expression
			call, ok := exprNode.X.(*ast.CallExpr)
			if !ok {
				return true
			}
			// assert the function is a selector expression
			selectExpr, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			if selectExpr.Sel.Name == "SetMaxOpenConns" {
				found = true
				// assert the argument in the call is a literal
				data, ok := call.Args[0].(*ast.BasicLit)
				if !ok {
					return true
				}
				if data.Value != "15" {
					pass.Reportf(selectExpr.Pos(), "SetMaxOpenConns with %s does not match required %s", data.Value, requiredVal)
				}

				return false // found it, stop recursing.
			}

			return true
		})
	}
	if !found {
		return nil, fmt.Errorf("not found")
	}
	return nil, nil
}
