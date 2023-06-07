/**
 * パッケージ名: repl
 * ファイル名: repl.go
 * 概要: REPLを実装する
 * REPLとは、Read(読み込み)-Eval(評価)-Print(表示)-Loop(繰り返し)の略で、対話型のプログラムを実行するための環境のこと。
 */
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MasaruFukazawa/monkey-lang/src/lexer"
	"github.com/MasaruFukazawa/monkey-lang/src/token"
)

const PROMPT = ">> "

/**
 * 関数名: Start
 * 処理: REPLを開始する
 * 引数: 入力, 出力
 * 戻値: なし
 */
func Start(in io.Reader, out io.Writer) {

	// 入力を取変
	// .. bufioパッケージのScanner構造体を使う
	scanner := bufio.NewScanner(in)

	// プロンプトを表示
	fmt.Printf(PROMPT)

	for scanner.Scan() {

		// 入力された文字列を取変
		// .. 1行ずつ読み込む
		line := scanner.Text()

		// 入力された文字列を字句解析する
		// .. 1行ずつ字句解析する
		l := lexer.New(line)

		// 字句解析を行う
		// .. トークンを1つずつ読み込む
		// .. EOFに達したら続ばない
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// トークンを表示
			fmt.Printf("%+v\n", tok)
		}

		// プロンプトを表示
		fmt.Printf(PROMPT)
	}
}
