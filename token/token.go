package token

// トークンを定義する
const (
	ILLEGAL   = "ILLEGAL" // 規則違反
	EOF       = "EOF"     // ファイルの終端

	// 識別子(変数名・関数名) : ユーザが宣言する名前
	IDENT     = "IDENT"

	// リテラル : 扱うデータの型
	INT       = "INT"

	// 演算子 : 使用できる演算子
	ASSIGN    = "="
	PLUS      = "+"

	// デリミタ(区切り文字) : コード上の区切り文字
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN    = "("
	RPAREN    = ")"

	LBRACE    = "{"
	RBRACE    = "}"

	// キーワード : コード上で使用する予約語
	FUNCTION  = "FUNCTION" // 関数定義
	LET       = "LET"      // 変数定義
)

type TokenType string

// トークンを表す構造体
type Token struct {
	Type    TokenType // トークンタイプ
	Literal string    // トークン文字列
}