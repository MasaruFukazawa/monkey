/**
 * パッケージ名: parser
 * ファイル名: parser.go
 * 概要: 構文解析器を実装する
 */
package parser

import (
	"fmt"

	"github.com/MasaruFukazawa/monkey-lang/src/ast"
	"github.com/MasaruFukazawa/monkey-lang/src/lexer"
	"github.com/MasaruFukazawa/monkey-lang/src/token"
)

// 構文解析器を表す構造体
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

/**
 * 名前: New
 * 処理: 構文解析器のポインタを返す
 * 引数: *lexer.Lexer
 * 戻値: *Parser
 */
func New(l *lexer.Lexer) *Parser {

	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

/**
 * 名前: Parser.nextToken
 * 処理: 次のトークンを読み込む
 * 引数: なし
 * 戻値: なし
 */
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/**
 * 名前: Parser.ParseProgram
 * 処続: 構文解析を行う
 * 引数: なし
 * 戻値: *ast.Program
 */
func (p *Parser) ParseProgram() *ast.Program {

	// programにast.Programのポインタを代入
	program := &ast.Program{}

	// program.Statementsにast.Statementを追加していく
	program.Statements = []ast.Statement{}

	// EOFに達するまで繰り返す
	for p.curToken.Type != token.EOF {

		// statementに構文解析結果を代入
		stmt := p.parseStatement()

		// stmtがnilでなければ、program.Statementsに追加
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		// 次のトークンへ進める
		p.nextToken()

	} // p.nextToken()

	return program
}

/**
 * 名前: Parser.parseStatement
 * 処続: 構文解析を行う
 * 引数: なし
 * 戻値: ast.Statement
 */
func (p *Parser) parseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

/**
 * 名前: Parser.parseLetStatement
 * 処続: 構文解析を行う
 * 引数: なし
 * 戻値: *ast.LetStatement
 */
func (p *Parser) parseLetStatement() *ast.LetStatement {

	// letを持つast.LetStatementのポインタを生成
	stmt := &ast.LetStatement{Token: p.curToken}

	// 次のトークンがIDENTでなければnilを返す
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// IDENTを持つast.Identifierを生成
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 次のトークンがASSIGNでなければnilを返す
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: セミコロンに遭遇するまで式を読み飛ばしている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/**
 * 名前: Parser.parseReturnStatement
 * 処理: 構文解析を行う
 * 引数: なし
 * 戻値: *ast.ReturnStatement
 */
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {

	// returnを持つast.ReturnStatementのポインタを生成
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// 次のトークンへ進める
	p.nextToken()

	// TODO: セミコロンに遭遇するまで式を読み飛ばしている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/**
 * 名前: Parser.curTokenIs
 * 処続: 現在のトークンが引数のトークンと同じかどうかを返す
 * 引数: token.TokenType
 * 戻値: bool
 */
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

/**
 * 名前: Parser.peekTokenIs
 * 処続: 次のトークンが引数のトークンと同じかどうかを返す
 * 引数: token.TokenType
 * 戻値: bool
 */
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

/**
 * 名前: Parser.expectPeek
 * 処続: 次のトークンが引数のトークンと同じであれば、次のトークンへ進める
 * 引数: token.TokenType
 * 戻値: bool
 */
func (p *Parser) expectPeek(t token.TokenType) bool {

	// 次のトークンが引数のトークンと同じであれば、次のトークンへ進める
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		// 次のトークンが引数のトークンと同じでなければ、エラーを追加
		p.peekError(t)
		return false
	}
}

/**
 * 名前: Parser.errors
 * 処続: 構文解析中に発生したエラーを返す
 * 引数: なし
 * 戻値: []string
 */
func (p *Parser) Errors() []string {
	return p.errors
}

/**
 * 名前: Parser.peekError
 * 処続: 次のトークンが引数のトークンと同じでなければ、エラーを追加する
 * 引数: token.TokenType
 * 戻値: なし
 */
func (p *Parser) peekError(t token.TokenType) {

	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)

	// エラーを追加
	p.errors = append(p.errors, msg)
}
