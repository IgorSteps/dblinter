package domain

import (
	"go/ast"
	"go/token"
	"go/types"
)

type CallSite struct {
	Receiver types.Type
	Method   string
	Args     []ast.Expr
	Position token.Pos
}
