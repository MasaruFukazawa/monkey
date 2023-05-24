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
  * 名前: newToken
  * 処理: トークン構造体を生成する
  * 引数: トークンの種類, トークンの文字
  * 戻値: トークン構造体
  */
  func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
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

	// 空白文字を読み飛ばす
	l.skipWhitespace()

	// 現在検査中の文字に応じてトークンを返す
	// .. default 以外は、1文字で意味が完結するトークン
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

		// 英字でない場合、規則違反とする
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
	for isLetter(l.ch) {
		l.readChar()
	}

	// 識別子(変数名・関数名)を返す
	// .. 英字のみの文字列を返す
	return l.input[position:l.position]
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
	for isDigit(l.ch) {
		l.readChar()
	}

	// 整数を返す
	// .. 数字のみの文字列を返す
	return l.input[position:l.position]
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