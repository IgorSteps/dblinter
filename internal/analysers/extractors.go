package analysers

import (
	"go/ast"

	"github.com/IgorSteps/dblinter/internal/rules"
	"golang.org/x/tools/go/analysis"
)

func FindCallsSites(pass *analysis.Pass) []rules.CallSite {
	callSites := []rules.CallSite{}
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// A call is an ExprStmt as such confirm node is indeed an ExprStmt.
			exprNode, ok := node.(*ast.ExprStmt)
			if !ok {
				return true
			}

			// Assert that the node is a function call expression.
			call, ok := exprNode.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Assert the call is a function ie., selector expression.
			selectExpr, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			callSite := rules.CallSite{
				Receiver: pass.TypesInfo.TypeOf(selectExpr.X).String(),
				Method:   selectExpr.Sel.Name,
				Args:     call.Args,
				Position: call.Pos(),
			}
			callSites = append(callSites, callSite)

			return true
		})
	}
	return callSites
}
