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

	"github.com/MasaruFukazawa/monkey-lang/src/evaluator"
	"github.com/MasaruFukazawa/monkey-lang/src/lexer"
	"github.com/MasaruFukazawa/monkey-lang/src/object"
	"github.com/MasaruFukazawa/monkey-lang/src/parser"
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
	env := object.NewEnvironment()

	for {

		// プロンプトを表示
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		// 入力された文字列を取変
		// .. 1行ずつ読み込む
		line := scanner.Text()

		// 入力された文字列を字句解析する
		// .. 1行ずつ字句解析する
		l := lexer.New(line)

		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {

	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
