/**
 * パッケージ名: parser
 * ファイル名: parser_test.go
 * 概要: parserのテストを実装する
 */
package parser

import (
	"testing"

	"github.com/MasaruFukazawa/monkey-lang/src/ast"
	"github.com/MasaruFukazawa/monkey-lang/src/lexer"
)

func TestLetStatements(t *testing.T) {

	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	// パーサーのエラーをチェック
	checkParserErrors(t, p)

	// programがnilでないことを確認
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	// program.Statementsの長さが3でないことを確認
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

/**
 * 名前: testLetStatement
 * 概要: LetStatementのテストを実装する
 * 引数: t *testing.T, s Statement, name string
 * 戻り値: bool
 */
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {

	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

/**
 * 名前: TestReturnStatements
 * 概要: return文のテストを実装する
 * 引数: t *testing.T
 * 戻り値:
 */
func TestReturnStatements(t *testing.T) {

	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	// program.Statementsの長さが3でないことを確認
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {

		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.returnStatement, got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {

			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())

		}
	}
}

/**
 * 名前: checkParserErrors
 * 処理: パーサーのエラーをチェックする
 * 引数:
 * .. t *testing.T: テスト
 * .. p *Parser: パーサー
 * 戻り値:
 */
func checkParserErrors(t *testing.T, p *Parser) {

	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
