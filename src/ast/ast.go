/**
 * パッケージ名: ast
 * ファイル名: ast.go
 * 概要: 抽象構文木の定義
 */
package ast

import (
	"bytes"
	"github.com/MasaruFukazawa/monkey-lang/src/token"
	"strings"
)

// 抽象構文木のノードのインターフェース
type Node interface {
	// Nodeを継承する構造体は、TokenLiteral()メソッドを実装しなければならない
	TokenLiteral() string

	// デバック用に抽象構文木を文字列にして返す
	// Nodeを継承する構造体は、String()メソッドを実装しなければならない
	String() string
}

// 抽象構文木の「文」のインターフェース
type Statement interface {
	// Nodeを継承する構造体は、TokenLiteral()メソッドを実装しなければならない
	Node

	// Statementを継承する構造体は、statementNode()メソッドを実装しなければならない
	statementNode()
}

// 抽象構文木の「式」のインターフェース
type Expression interface {
	// Nodeを継承する構造体は、TokenLiteral()メソッドを実装しなければならない
	Node

	// Expressionを継承する構造以件のexpressionNode()メタッドを実装してない
	expressionNode()
}

// LET文を表すノード
// .. Statementインターフェースを満たす
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 変数名
	Value Expression  // 変数名にバインドする式
}

/**
 * 名前: LetStatement.statementNode
 * 概要:
 *	LET文のトークンリテラルを返す
 *  Statementインターフェースを満たす
 */
func (ls *LetStatement) statementNode() {}

/**
 * 名前: LetStatement.TokenLiteral
 * 概要:
 *	LET文のトークンリテラルを返す
 *  Statementインターフェースを満たす
 *  .. Nodeインターフェースを満たす
 */
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

/**
 * 名前: LetStatement.String
 * 概要:
 *	LET文のトークンリテラルを返す
 *  Statementインターフェースを満たす
 *  .. Nodeインターフェースを満たす
 */
func (ls *LetStatement) String() string {

	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	// Valueがnilでない場合
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Return文を表すノード
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression  // return文の返り値
}

/**
 * 名前: ReturnStatement.statementNode
 * 概要:
 *	Return文のトークンリテラルを返す
 *  Statementインターフェースを満たす
 */
func (rs *ReturnStatement) statementNode() {}

/**
 * 名前: ReturnStatement.TokenLiteral
 * 概要:
 *	Return文のトークンリテラルを返す
 *  TokenLiteralインターフェースを満たす
 */
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

/**
 * 名前: ReturnStatement.String
 * 概要:
 *	Return文のトークンリテラルを返す
 *  Nodeインターフェースを満たす
 */
func (rs *ReturnStatement) String() string {

	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	// ReturnValueがnilでない場合
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

/**
 * 名前: ExpressionStatement
 * 概要:
 *	式のノード
 */
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression  // 式を保持するフィールド
}

/**
 * 名前: ExpressionStatement.statementNode
 * 概要:
 *	ExpressionStatementのトークンリテラルを返す
 *  Statementインターフェースを満たす
 */
func (es *ExpressionStatement) statementNode() {}

/**
 * 名前: ExpressionStatement.TokenLiteral
 * 概要:
 *	ExpressionStatementのトークンリテラルを返す
 *  TokenLiteralインターフェースを満たす
 */
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

/**
 * 名前: ExpressionStatement.String
 * 概要:
 *	ExpressionStatementのトークンリテラルを返す
 *  Nodeインターフェースを満たす
 */
func (es *ExpressionStatement) String() string {

	// Expressionがnilの場合
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// 識別子(変数名・関数名)を表すノード
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string      // 変数名
}

/**
 * 名前: Identifier.expressionNode
 * 概要:
 * 	識別子(変数名・関数名)のトークンリテラルを返す
 *	Expressionインターフェースを満たす
 */
func (i *Identifier) expressionNode() {}

/**
 * 名前: Identifier.TokenLiteral
 * 概要:
 *	識別子(変数名・関数名)のトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

/**
 * 名前: Identifier.String
 * 概要:
 *	識別子(変数名・関数名)のトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (i *Identifier) String() string {
	return i.Value
}

// 整数リテラルを表すノード
type IntegerLiteral struct {
	Token token.Token // token.INT トークン
	Value int64       // 整数リテラルの値
}

/**
 * 名前: IntegerLiteral.expressionNode
 * 概要:
 *	整数リテラルのトークンリテラルを返す
 *	Expressionインターフェースを満たす
 */
func (il *IntegerLiteral) expressionNode() {}

/**
 * 名前: IntegerLiteral.TokenLiteral
 * 概要:
 *	整数リテラルのトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

/**
 * 名前: IntegerLiteral.String
 * 概要:
 *	整数リテラルのトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// 前置演算子を表すノード
type PrefixExpression struct {
	Token    token.Token // 前置演算子トークン、例えば「!」
	Operator string      // 前置演算子、例えば「-」
	Right    Expression  // 右側の式
}

/**
 * 名前: PrefixExpression.expressionNode
 * 概要:
 *	前置演算子のトークンリテラルを返す
 *	Expressionインターフェースを満たす
 */
func (pe *PrefixExpression) expressionNode() {}

/**
 * 名前: PrefixExpression.TokenLiteral
 * 概要:
 *	前置演算子のトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

/**
 * 名前: PrefixExpression.String
 * 概要:
 *	前置演算子のトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (pe *PrefixExpression) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// 中置演算子を表すノード
type InfixExpression struct {
	Token    token.Token // 演算子トークン、例えば「+」
	Left     Expression  // 左側の式
	Operator string      // 演算子
	Right    Expression  // 右側の式
}

/**
 * 名前: InfixExpression.expressionNode
 * 概要:
 *	中置演算子のトークンリテラルを返す
 *	Expressionインターフェースを満たす
 */
func (oe *InfixExpression) expressionNode() {}

/**
 * 名前: InfixExpression.TokenLiteral
 * 概要:
 *	中置演算子のトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

/**
 * 名前: InfixExpression.String
 * 概要:
 *	中置演算子のトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (oe *InfixExpression) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

/**
 *
 * 真偽値を表すノード
 *
 */
type Boolean struct {
	Token token.Token
	Value bool
}

/**
 * 名前: Boolean.expressionNode
 * 概要:
 *	真偽値のトークンリテラルを返す
 *	Expressionインターフェースを満たす
 */
func (b *Boolean) expressionNode() {}

/**
 * 名前: Boolean.TokenLiteral
 * 概要:
 *	真偽値のトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

/**
 * 名前: Boolean.String
 * 概要:
 *	真偽値のトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (b *Boolean) String() string {
	return b.Token.Literal
}

/**
 *
 * if文を表すノード
 *
 */
type IfExpression struct {
	Token       token.Token     // 'if' トークン
	Condition   Expression      // 条件式
	Consequence *BlockStatement // 条件が真の場合の文
	Alternative *BlockStatement // 条件が偽の場合の文
}

/**
 * 名前: IfExpression.expressionNode
 * 概要:
 *	if文のトークンリテラルを返す
 *	Expressionインターフェースを満たす
 */
func (ie *IfExpression) expressionNode() {}

/**
 * 名前: IfExpression.TokenLiteral
 * 概要:
 *	if文のトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

/**
 * 名前: IfExpression.String
 * 概要:
 *	if文のトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (ie *IfExpression) String() string {

	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

/**
 * 名前: BlockStatement
 * 概要:
 *	ブロック文を表すノード
 *	Statementインターフェースを満たす
 */
type BlockStatement struct {
	Token      token.Token // '{' トークン
	Statements []Statement // ブロック文の中の文
}

/**
 * 名前: BlockStatement.statementNode
 * 概要:
 *	BlockStatementのトークンリテラルを返す
 *	Statementインターフェースを満たす
 */
func (bs *BlockStatement) statementNode() {}

/**
 * 名前: BlockStatement.TokenLiteral
 * 概要:
 *	BlockStatementのトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

/**
 * 名前: BlockStatement.String
 * 概要:
 *	BlockStatementのトークンリテラルを返す
 *	Nodeインターフェースを満たす
 */
func (bs *BlockStatement) String() string {

	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/**
 *
 * 名前: 関数リテラルを表すノード
 *
 */
type FunctionLiteral struct {
	Token      token.Token     // 'fn' トークン
	Parameters []*Identifier   // パラメータリスト
	Body       *BlockStatement // 関数の本体
}

/**
 * 名前: FunctionLiteral.expressionNode
 * 概要:
 *  関数リテラルのトークンリテラルを返す
 * 	Expressionインターフェースを満たす
 */
func (fl *FunctionLiteral) expressionNode() {}

/**
 * 名前: FunctionLiteral.TokenLiteral
 * 概要:
 *  関数リテラルのトークンリテラルを返す
 *	TokenLiteralインターフェースを満たす
 */
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

/**
 * 名前: FunctionLiteral.String
 * 概要:
 *  関数リテラルのトークンリテラルを返す
 *  Nodeインターフェースを満たす
 */
func (fl *FunctionLiteral) String() string {

	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// プログラム全体を表すノード
// .. Nodeインターフェースを満たす
type Program struct {

	// プログラム全体の文の配列
	Statements []Statement
}

/**
 * 名前: TokenLiteral
 * 概要: プログラム全体の文の配列の先頭のトークンリテラルを返す
 */
func (p *Program) TokenLiteral() string {

	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		// 空のプログラムの場合は空文字列を返す
		return ""
	}
}

/**
 * 名前: String
 * 概要: デバック用に抽象構文木を文字列にして返す
 */
func (p *Program) String() string {

	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
