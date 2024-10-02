/**
 * パッケージ名: lexer
 * ファイル名: lexer.go
 * 概要: 字句解析器を実装する
 * 字句解析とは、文字列をトークンに分割する処理のこと。
 */
package lexer

import (
	"github.com/MasaruFukazawa/monkey-lang/src/token"
)

// 字句解析器を表す構造体
type Lexer struct {
	input        string // ソースコード文字列
	position     int    // 入力における現在の位置 : 現在の文字を指し示す。 初期値は0
	readPosition int    // これから読み込む位置 : 現在の文字の次を指し示す。初期値は0
	ch           byte   // 現在検査中の1文字
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

	// position（現在の読み込み位置）とreadPosition（次に読み込む位置）を1つ進める
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

	// 空白文字を読み飛ばす
	l.skipWhitespace()

	// 現在検査中の文字に応じてトークンを返す
	// .. default 以外は、1文字で意味が完結するトークン
	switch l.ch {
	case '=':
		// 1文字前を覗き見する
		if l.peekChar() == '=' { // == であれば、EQトークンとする
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else { // = であれば、ASSIGNトークンとする
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// 1文字前を覗き見する
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0: // ソースコードの終端に達した場合
		tok.Literal = ""
		tok.Type = token.EOF

	// ユーザ定義の識別子(変数名・関数名)を読み込む
	default:

		// 文字が英字である限り、識別子として読み込む
		if isLetter(l.ch) {

			// 識別子(変数名・関数名)を取得する
			tok.Literal = l.readIdentifier()

			// 識別子(変数名・関数名)の種類を判定する
			tok.Type = token.LookupIdent(tok.Literal)

			return tok

			// 文字が数字である限り、整数として読み込む
		} else if isDigit(l.ch) {

			// 整数を取得する
			tok.Type = token.INT
			tok.Literal = l.readNumber()

			return tok

			// 英字のみでもないし、数字のみでもない場合、ILLEGALトークンとする
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	// 1文字読み込む
	l.readChar()

	return tok
}

/**
 * 名前: readIdentifier
 * 処理: 識別子(変数名・関数名)を読み込む
 * 引数: なし
 * 戻値: 識別子(変数名・関数名)
 */
func (l *Lexer) readIdentifier() string {

	// 識別子(変数名・関数名)の開始位置を記憶
	position := l.position

	// 英字である限り、1文字ずつ読み込む
	// .. l.positionを1つ進める
	// .. l.readPositionを1つ進める
	// .. 英字である限り、1文字ずつ読み込む
	// .. なので、識別子(変数名・関数名)の終了位置は、l.positionの1つ前になる
	// .. l.positionは、Letterでない文字を指し示す
	for isLetter(l.ch) {
		l.readChar()
	}

	// 識別子(変数名・関数名)を返す
	// .. 英字のみの文字列を返す
	// .. positionからl.positionまでの文字列を返す
	return l.input[position:l.position]
}

/**
 * 名前: skipWhitespace
 * 処理: 空白文字を読み飛ばす
 * 引数: なし
 * 戻値: なし
 */
func (l *Lexer) skipWhitespace() {

	// 空白文字である限り、1文字ずつ読み込む
	// .. l.positionを1つ進める
	// .. l.readPositionを1つ進める
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}

/**
 * 名前: readNumber
 * 処理: 整数を読み込む
 * 引数: なし
 * 戻値: 整数
 */
func (l *Lexer) readNumber() string {

	// 整数の開始位置を記憶
	position := l.position

	// 数字である限り、1文字ずつ読み込む
	// .. l.positionを1つ進める
	// .. l.readPositionを1つ進める
	// .. 数字である限り、1文字ずつ読み込む
	// .. なので、整数の終了位置は、l.positionの1つ前になる
	// .. l.positionは、数字でない文字を指し示す
	for isDigit(l.ch) {
		l.readChar()
	}

	// 整数を返す
	// .. 数字のみの文字列を返す
	// .. positionからl.positionまでの文字列を返す
	return l.input[position:l.position]
}

/**
 * 名前: peekChar
 * 処理: 次の文字を読み込む。ただし、読み込み位置は進めない
 * 引数: なし
 * 戻り値: なし
 */
func (l *Lexer) peekChar() byte {

	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}

}

/**
 * 名前: readString
 * 処理: 文字列を読み込む
 * 引数: なし
 * 戻値: 文字列
 */
func (l *Lexer) readString() string {

	position := l.position + 1

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

/**
 * 名前: lexer.New
 * 処理: Lexer構造体のポインタを返す
 * 引数: input : ソースコード文字列
 * 戻値: *Lexer
 */
func New(input string) *Lexer {

	// lexer構造体のポインタを返す
	l := &Lexer{input: input}

	// 1文字読み込む
	// .. l.ch = l.input[0]
	// .. l.position = 0
	// .. l.readPosition = 1
	l.readChar()

	return l
}

/**
 * 名前: newToken
 * 処理: トークンごとにトークン構造体を生成する
 * 引数: トークンの種類, トークンの文字
 * 戻値: トークン構造体
 */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/**
 * 名前: isLetter
 * 処理: 文字が英字かどうかを判定する
 * 引数: 文字
 * 戻値: bool
 */
func isLetter(ch byte) bool {
	// 英字（大文字小文字）、_ であればtrueを返す
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/**
 * 名前: isDigit
 * 処理: 文字が数字かどうかを判定する
 * 引数: 文字
 * 戻値: bool
 */
func isDigit(ch byte) bool {
	// 数字であればtrueを返す
	return '0' <= ch && ch <= '9'
}
