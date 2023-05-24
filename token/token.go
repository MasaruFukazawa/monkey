package token


// トークンの種類を定義する
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


// 予約語のマップ
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
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