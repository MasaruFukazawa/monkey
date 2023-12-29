/**
 * パッケージ名: parser
 * ファイル名: parser_test.go
 * 概要: parserのテストを実装する
 */
package parser

import (
	"fmt"
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
 * 名前: TestIdentifierExpression
 * 概要: 識別子式のテストを実装する
 * 引数: t *testing.T
 * 戻り値:
 */
func TestIdentifierExpression(t *testing.T) {

	input := "foobar;"

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	// program.Statementsの長さが1でないことを確認
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	// program.Statements[0]がast.ExpressionStatementであることを確認
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

/*
 * 名前: TestIntegerLiteralExpression
 * 処理: 整数リテラル式のテストを実装する
 * 引数: t *testing.T
 * 戻り値:
 */
func TestIntegerLiteralExpression(t *testing.T) {

	input := "5;"

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	// program.Statementsの長さが1でないことを確認
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	// program.Statements[0]がast.ExpressionStatementであることを確認
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

/**
 * 名前: TestParsingPrefixExpressions
 * 概要: 前置式の解析テストを実装する
 * 引数: t *testing.T
 * 戻り値:
 */
func TestParsingPrefixExpressions(t *testing.T) {

	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {

		l := lexer.New(tt.input)

		p := New(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d¥n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statement[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

/**
 * 名前: TestParsingInfixExpressions
 * 概要: 中置式の解析テストを実装する
 * 引数: t *testing.T
 * 戻り値:
 */
func TestParsingInfixExpressions(t *testing.T) {

	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {

		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d¥n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

/**
 * 名前: TestOperatorPrecedenceParsing
 * 概要: 演算子の優先順位の解析テストを実装する
 * 引数: t *testing.T
 * 戻り値:
 */
func TestOperatorPrecedenceParsing(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
	}

	for _, tt := range tests {

		l := lexer.New(tt.input)

		p := New(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}

	}

}

/*
 * 名前: testIntegerLiteral
 * 概要: 整数リテラルのテストを実装する
 * 引数:
 * .. t *testing.T
 * .. il ast.Expression: 式
 * .. value int64: 整数値
 * 戻り値:
 * .. bool
 */
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {

	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

/*
 * 名前: checkParserErrors
 * 処理: パーサーのエラーをチェックする
 * 引数:
 * .. t *testing.T: テスト
 * .. p *Parser: パーサー
 * 戻り値:
 * .. なし
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
