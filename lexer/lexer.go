package lexer


import (
	"github.com/MasaruFukazawa/monkey-lang/token"
)


type Lexer struct {
	input string
	position int     // 入力における現在の位置 : 現在の文字を指し示す
	readPosition int // これから読み込む位置 : 現在の文字の次を指し示す
	ch byte          // 現在検査中の文字
}


/**
 * 名前: New
 * 処理: Lexer構造体のポインタを返す
 * 引数: input : ソースコード文字列
 * 戻値: *Lexer
 */ 
func New(input string) *Lexer {

	// lexer構造体のポインタを返す
	l := &Lexer{input: input}

	// 1文字読み込む
	l.readChar()

	return l
}


/**
 * 関数名: readChar
 * 処理: 1文字読み込む
 * 引数: なし
 * 戻値: なし
 */
func (l *Lexer) readChar() {

	// 入力が終端に達しているかどうかを検査
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// chに次の文字を代入
		l.ch = l.input[l.readPosition]
	}

	// positionとreadPositionを1つ進める
	l.position = l.readPosition
	l.readPosition += 1
}


/**
 * 関数名: NextToken
 * 処理: 1文字読み込み、その文字のトークン構造体データを返す
 * 引数: なし
 * 戻値: トークン構造体データ
 */
 func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	// 現在検査中の文字に応じてトークンを返す
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// 1文字読み込む
	l.readChar()

	return tok
 }


 /**
  * 名前: newToken
  * 処理: トークン構造体を生成する
  * 引数: トークンの種類, トークンの文字
  * 戻値: トークン構造体
  */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}	