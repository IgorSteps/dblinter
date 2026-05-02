package domain

import (
	"go/ast"
	"go/types"
)

type CallSite struct {
	Receiver types.Type
	Method   string
	Args     []ast.Expr
}
