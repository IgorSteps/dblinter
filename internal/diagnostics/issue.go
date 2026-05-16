package diagnostics

import "go/token"

type Issue struct {
	RuleID  string
	Message string
	Doc     string
	Pos     token.Pos
}
