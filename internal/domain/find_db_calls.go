package domain

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func FindDBCalls(pass *analysis.Pass) []CallSite {
	callSites := []CallSite{}
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
				callSite := CallSite{
					Receiver: pass.TypesInfo.TypeOf(selectExpr.X),
					Method:   selectExpr.Sel.Name,
					Args:     call.Args,
				}
				callSites = append(callSites, callSite)
			}

			return true
		})
	}
	return callSites
}
