package lexer

import (
	"testing"
	"github.com/MasaruFukazawa/monkey-lang/token"
)


func TestNextToken(t *testing.T) {

	// テストデータとなる文字列を定義
	input := `let five = 5;
let ten = 10;
	
let add = fn(x, y) {
	x + y;
};
		
let result = add(five, ten);`

	// テスト結果となるトークンの期待値を定義
	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		// let five = 5;
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		/*
		let add = fn(x, y) {
			x + y;
		};
		*/
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// ファイルの終端
		{token.EOF, ""},
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

