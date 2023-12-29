/**
 * パッケージ名: ast
 * ファイル名: ast_test.go
 * 概要: astのテストを実装する
 */
package ast

import (
	"testing"

	"github.com/MasaruFukazawa/monkey-lang/src/token"
)

func TestString(t *testing.T) {

	program := &Program{
		Statements: []Statement{
			// let myVar = anotherVar; の抽象構文木表現
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}

}
