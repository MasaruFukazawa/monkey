/**
 * パッケージ名: ast
 * ファイル名: ast.go
 * 概要: 抽象構文木の定義
 */
package ast

import (
	"github.com/MasaruFukazawa/monkey-lang/src/token"
)

// 抽象構文木のノードのインターフェース
type Node interface {
	// Nodeを継承する構造体は、TokenLiteral()メソッドを実装しなければならない
	TokenLiteral() string
}

// 抽象構文木の文のインターフェース
type Statement interface {
	// Nodeを継承する構造体は、TokenLiteral()メソッドを実装しなければならない
	Node
	// Statementを継承する構造体は、statementNode()メソッドを実装しなければならない
	statementNode()
}

// 抽象構文木の式のインターフェース
type Expression interface {
	// Nodeを継承する構造体は、TokenLiteral()メソッドを実装しなければならない
	Node
	// Expressionを継承する構造以件のexpressionNode()メタッドを実装してない
	expressionNode()
}

// LET文を表すノード
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
 *  TokenLiteralインターフェースを満たす
 */
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
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

// 識別子(変数名・関数名)を表すノード
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string      // 変数名
}

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

// プログラム全体を表すノード
type Program struct {

	// プログラム全体の文の配列
	Statements []Statement
}

/**
 * 名前: TokenLiteral
 * 概要:
 */
func (p *Program) TokenLiteral() string {

	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		// 空のプログラムの場合は空文字列を返す
		return ""
	}
}
