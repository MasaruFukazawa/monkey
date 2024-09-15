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

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalIntegerInfixExpression(node.Operator, left, right)
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

func evalPrefixExpression(operator string, right object.Object) object.Object {

	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	case "+":
		return evalPlusPrefixOperatorExpression(right)
	default:
		return NULL
	}

}

func evalBangOperatorExpression(right object.Object) object.Object {

	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}

}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {

	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value

	return &object.Integer{Value: -value}
}

func evalPlusPrefixOperatorExpression(right object.Object) object.Object {

	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value

	return &object.Integer{Value: value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {

	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	default:
		return NULL
	}

}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {

	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
	default:
		return NULL
	}

}
