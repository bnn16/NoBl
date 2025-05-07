package ast

import (
	"nobl/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ConstStatement{
				Token: token.Token{Type: token.CONST, Literal: "const"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "const myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
