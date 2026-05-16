package rules

import (
	"go/ast"
	"go/token"
)

type CallSite struct {
	Receiver string
	Method   string
	Args     []ast.Expr
	Position token.Pos
}
