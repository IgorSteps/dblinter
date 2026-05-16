package diagnostics

import "go/token"

type Issue struct {
	RuleID  string
	Message string
	Pos     token.Pos
}
