/**
 * パッケージ名: token
 * ファイル名: token.go
 * 概要: トークンを定義する
 * トークンとは、ソースコードを構成する要素のこと。
 */
package token

// monkeylangで使用できるトークンの種類を定義する
// .. トークンは、字句解析の結果として得られる
const (
	ILLEGAL = "ILLEGAL" // 規則違反
	EOF     = "EOF"     // ファイルの終端

	// 識別子(変数名・関数名) : ユーザが宣言する名前
	IDENT = "IDENT"

	// リテラル : 扱うデータの型
	INT    = "INT"
	STRING = "STRING"

	// 演算子 : 使用できる演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// 比較演算子 : 使用できる比較演算子
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// デリミタ(区切り文字) : コード上の区切り文字
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// キーワード : コード上で使用する予約語
	FUNCTION = "FUNCTION" // 関数定義
	LET      = "LET"      // 変数定義
	TRUE     = "TRUE"     // 真
	FALSE    = "FALSE"    // 偽
	IF       = "IF"       // 構文構造使用: 条件分岐
	ELSE     = "ELSE"     // 構文構造使用: 条件分岐
	RETURN   = "RETURN"   // 構文構造使用: 関数からの戻り値
)

type TokenType string

// トークンを表す構造体
type Token struct {
	Type    TokenType // トークンの種類
	Literal string    // トークン文字列（ 変数名 や + , - などの文字列 ）
}

// 予約語のマップ
// .. 予約語は、言語の構文構造に使用するキーワード
// .. 予約語は、変数名や関数名として使用できない
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

/**
 * 名前: LookupIdent
 * 処理: 予約語かどうかを判定する
 * 引数: ident: 識別子(変数名・関数名)の文字列
 * 戻り値: TokenType
 */
func LookupIdent(ident string) TokenType {

	// 予約語かどうかを判定する
	// .. kewordsに識別子(変数名・関数名)が存在する場合、予約語とする
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	// 予約語でない場合、識別子(変数名・関数名)とする
	return IDENT
}
