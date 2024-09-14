package evaluator

import (
	"github.com/MasaruFukazawa/monkey-lang/src/ast"
	"github.com/MasaruFukazawa/monkey-lang/src/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

/**
 * 関数名: Eval
 * 処理: 引数で渡された抽象構文木を評価する
 * 引数: 抽象構文木
 * 戻値: 評価結果
 */
func Eval(node ast.Node) object.Object {

	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBooleanObject(node.Value)
	}

	return nil
}

func evalStatements(stms []ast.Statement) object.Object {

	var result object.Object

	for _, statement := range stms {
		result = Eval(statement)
	}

	return result
}

func nativeBooleanObject(input bool) *object.Boolean {

	if input {
		return TRUE
	}

	return FALSE
}
