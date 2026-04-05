package domain

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type Analyser struct{}

func (s *Analyser) Run(pass *analysis.Pass, targets []*FunctionCall) ([]*FunctionCall, error) {
	found := false
	for _, target := range targets {
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

				if selectExpr.Sel.Name == target.Name {
					found = true
					// assert the argument in the call is a literal
					data, ok := call.Args[0].(*ast.BasicLit)
					if !ok {
						return true
					}
					target.Arg = data.Value
					return false // found it, stop recursing.
				}

				return true
			})
		}
	}

	if !found {
		return nil, fmt.Errorf("failed to find targets")
	}
	return targets, nil
}
