/**
 * パッケージ名: parser
 * ファイル名: parser.go
 * 概要: 構文解析器を実装する
 */
package parser

import (
	"fmt"
	"strconv"

	"github.com/MasaruFukazawa/monkey-lang/src/ast"
	"github.com/MasaruFukazawa/monkey-lang/src/lexer"
	"github.com/MasaruFukazawa/monkey-lang/src/token"
)

const (
	_ int = iota
	// LOWEST: 優先順位の最低値
	LOWEST
	// EQUALS: ==
	EQUALS
	// LESSGREATER: > または <
	LESSGREATER
	// SUM: +
	SUM
	// PRODUCT: *
	PRODUCT
	// PREFIX: -X または !X
	PREFIX
	// CALL: myFunction(X)
	CALL
)

// 優先順位のマップ
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,      // ==
	token.NOT_EQ:   EQUALS,      // !=
	token.LT:       LESSGREATER, // <
	token.GT:       LESSGREATER, // >
	token.PLUS:     SUM,         // +
	token.MINUS:    SUM,         // -
	token.SLASH:    PRODUCT,     // /
	token.ASTERISK: PRODUCT,     // *
	token.LPAREN:   CALL,        //
}

// 優先順位の定義
// 下の宣言を2つに分けると以下のようになる
// .. type prefixParseFn func() ast.Expression
// .. type infixParseFn func(ast.Expression) ast.Expression
type (
	// 前置構文解析関数
	prefixParseFn func() ast.Expression
	// 中置構文解析関数
	infixParseFn func(ast.Expression) ast.Expression
)

// 構文解析器を表す構造体
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	// 前置構文解析関数のマップ
	prefixParseFns map[token.TokenType]prefixParseFn

	// 中置構文解析関数のマップ
	infixParseFns map[token.TokenType]infixParseFn
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
	case token.LET: // let
		return p.parseLetStatement()
	case token.RETURN: // return
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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
 * 名前: Parser.parseExpressionStatement
 * 処続: 構文解析を行う
 * 引数: なし
 * 戻値: ast.ExpressionStatement
 */
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	// expressionを持つast.ExpressionStatementのポインタを生成
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	// 式を構文解析する
	stmt.Expression = p.parseExpression(LOWEST)

	// セミコロンに遭遇するまで式を読み飛ばす
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/**
 * 名前: Parse.noPrefixParseFnError
 * 処理: 構文解析を行う
 * 引数: token.TokenType
 * 戻値: なし
 */
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

/**
 * 名前: Parser.parseExpression
 * 処続: 構文解析を行う
 * 引数: int
 * 戻値: ast.Expression
 */
func (p *Parser) parseExpression(precedence int) ast.Expression {

	// 前置構文解析関数を取得
	prefix := p.prefixParseFns[p.curToken.Type]

	// 前置構文解析関数がnilであればnilを返す
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	// 前置構文解析関数を実行
	leftExp := prefix()

	// 次のトークンがセミコロンでないかつ、引数の優先順位よりも低い場合は繰り返す
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {

		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

/**
 * 名前: Parser.parseIdentifier
 * 処続: 構文解析を行う
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parseIdentifier() ast.Expression {
	// IDENTを持つast.Identifierポインタを生成
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
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

/**
 * 名前: Parser.registerPrefix
 * 概要: 引数のトークンタイプに対応する前置構文解析関数を登録する
 * 引数: token.TokenType, func() ast.Expression
 * 戻値: なし
 */
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

/**
 * 名前: Parser.registerInfix
 * 概要: 引数のトークンタイプに対応する中置構文解析関数を登録する
 * 引数: token.TokenType, func(ast.Expression) ast.Expression
 * 戻値: なし
 */
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

/**
 * 名前: Parser.parseIntegerLiteral
 * 概要: 整数リテラルを構文解析する
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parseIntegerLiteral() ast.Expression {

	lit := &ast.IntegerLiteral{Token: p.curToken}

	// 文字列をint64に変換
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	// エラーが発生した場合はエラーを追加
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)

		// nilを返す
		return nil
	}

	// 整数リテラルの値を設定
	lit.Value = value

	return lit
}

/**
 * 名前: Parser.parsePrefixExpression
 * 概要: 前置演算子を構文解析する
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parsePrefixExpression() ast.Expression {

	// 前置演算子を持つast.PrefixExpressionポインタを生成
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	// 次のトークンへ進める
	p.nextToken()

	// 式を構文解析
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

/**
 * 名前: Parser.peekPrecedence
 * 概要: 次のトークンの優先順位を返す
 * 引数: なし
 * 戻値: int
 */
func (p *Parser) peekPrecedence() int {

	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	// トークンが見つからなかった場合はLOWESTを返す
	return LOWEST
}

/**
 * 名前: Parser.curPrecedence
 * 概要: 現在のトークンの優先順位を返す
 * 引数: なし
 * 戻値: int
 */
func (p *Parser) curPrecedence() int {

	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	// トークンが見つからなかった場合はLOWESTを返す
	return LOWEST
}

/**
 * 名前: Parser.parseInfixExpression
 * 概要: 中置演算子を構文解析する
 * 引数: ast.Expression
 * 戻値: ast.Expression
 */
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {

	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	// 現在のトークンの優先順位を取得
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

/**
 * 名前: Parser.parseBoolean
 * 概要: 真偽値を構文解析する
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

/**
 * 名前: Parser.parseGroupedExpression
 * 概要:
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parseGroupedExpression() ast.Expression {

	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

/**
 * 名前: Parser.parseIfExpression
 * 概要: if文を構文解析する
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parseIfExpression() ast.Expression {

	expression := &ast.IfExpression{Token: p.curToken}

	// 次のトークンへ進める
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// 条件式を構文解析
	p.nextToken()

	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// Consequenceを構文解析
	expression.Consequence = p.parseBlockStatement()

	// 次のトークンがELSEであれば、ELSEブロックを構文解析
	if p.peekTokenIs(token.ELSE) {

		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		// Alternativeを構文解析
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

/**
 * 名前: Parser.parseBlockStatement
 * 概要: ブロック文を構文解析する
 * 引数: なし
 * 戻値: *ast.BlockStatement
 */
func (p *Parser) parseBlockStatement() *ast.BlockStatement {

	block := &ast.BlockStatement{Token: p.curToken}

	// Statementsにast.Statementを追加していく
	block.Statements = []ast.Statement{}

	// 次のトークンへ進める
	p.nextToken()

	// 次のトークンがRBRACEでない場合は繰り返す
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {

		// Statementを構文解析
		stmt := p.parseStatement()

		// stmtがnilでなければ、block.Statementsに追加
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		// 次のトークンへ進める
		p.nextToken()
	}

	return block
}

/**
 * 名前: Parser.parseFunctionLiteral
 * 概要: 関数リテラルを構文解析する
 * 引数: なし
 * 戻値: ast.Expression
 */
func (p *Parser) parseFunctionLiteral() ast.Expression {

	// 関数リテラルを持つast.FunctionLiteralポインタを生成
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// 次のトークンがLPARENでなければnilを返す
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// 関数のパラメータを構文解析
	lit.Parameters = p.parseFunctionParameters()

	// 次のトークンがLBRACEでなければnilを返す
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// 関数の本体を構文解析
	lit.Body = p.parseBlockStatement()

	return lit
}

/**
 * 名前: Parser.parseFunctionParameters
 * 概要: 関数のパラメータを構文解析する
 * 引数: なし
 * 戻値: []*ast.Identifier
 */
func (p *Parser) parseFunctionParameters() []*ast.Identifier {

	// パラメータリスト
	identifiers := []*ast.Identifier{}

	// 次のトークンがRPARENであれば、nilを返す
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	// 次のトークンへ進める
	p.nextToken()

	// IDENTを持つast.Identifierを生成
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	// 次のトークンがCOMMAであれば、繰り返す
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	// 次のトークンがRPARENでなければnilを返す
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

/**
 * 名前: Parser.parseCallExpression
 * 概要: 呼び出し式を構文解析する
 * 引数: ast.Expression
 * 戻値: ast.Expression
 */
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {

	exp := &ast.CallExpression{Token: p.curToken, Function: function}

	exp.Arguments = p.parseCallArguments()

	return exp
}

/**
 * 名前: Parser.parseCallArguments
 * 概要: 引数リストを構文解析する
 * 引数: なし
 * 戻値: []ast.Expression
 */
func (p *Parser) parseCallArguments() []ast.Expression {

	args := []ast.Expression{}

	// 次のトークンがRPARENであれば、nilを返す
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	// 次のトークンへ進める
	p.nextToken()

	// 式を構文解析
	args = append(args, p.parseExpression(LOWEST))

	// 次のトークンがCOMMAであれば、繰り返す
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	// 次のトークンがRPARENでなければnilを返す
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args

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

	// 前置構文解析関数のマップを初期化
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)

	// 前置構文解析関数のマップに関数を登録
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	// 前置構文解析関数のマップに関数を登録
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	// BANGトークンを前置構文解析関数のマップに登録
	p.registerPrefix(token.BANG, p.parsePrefixExpression)

	// MINUSトークンを前置構文解析関数のマップに登録
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// 真偽値を前置構文解析関数のマップに登録
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)

	// LPARENトークンを前置構文解析関数のマップに登録
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	// if文の構文解析
	p.registerPrefix(token.IF, p.parseIfExpression)

	// fn (関数リテラル)の構文解析
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)

	// 中間構文解析関数のマップを初期化
	p.infixParseFns = make(map[token.TokenType]infixParseFn)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)

	p.nextToken()
	p.nextToken()

	return p
}
