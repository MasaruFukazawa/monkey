package lexer

import (
	"testing"
	"github.com/MasaruFukazawa/monkey-lang/token"
)


func TestextNextToken(t *testing.T) {

	// テストデータとなる文字列を定義
	input := `=+(){},;`

	// テスト結果となるトークンの期待値を定義
	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {

		_token := l.NextToken()

		if _token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, _token.Type)
		}

		if _token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, _token.Literal)
		}

	}
}

