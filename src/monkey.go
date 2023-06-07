package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MasaruFukazawa/monkey-lang/src/repl"
)

func main() {

	// ユーザー名を取得
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	// プロンプトを表示
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	// REPLを開始する
	repl.Start(os.Stdin, os.Stdout)
}
