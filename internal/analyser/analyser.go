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
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// db.SetMaxOpenWhatever is a ExprStmt
			// as such confirm node is indeed an ExprStmt
			exprNode, ok := node.(*ast.ExprStmt)
			if !ok {
				fmt.Println("node not an expr stmt")
				return true
			}
			fmt.Printf("expr stmt node %v", node)
			// assert that the node is a function call expression
			call, ok := exprNode.X.(*ast.CallExpr)
			if !ok {
				return true
			}
			fmt.Printf("call node %v", call)
			// assert the function is a selector expression
			selectExpr, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			fmt.Printf("selector expression: %s", selectExpr.Sel.Name)
			if selectExpr.Sel.Name == "SetMaxOpenConns" {
				pass.Reportf(selectExpr.Pos(), "found SetMaxOpenConns")
				return false // found it, stop recursing.
			}

			return true
		})
	}
	return nil, nil
}
