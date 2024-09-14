/**
 * パッケージ名: evaluator
 * ファイル名: evaluator_test.go
 * 概要: 評価器のテストを実装する
 */
package evaluator

import (
	"github.com/MasaruFukazawa/monkey-lang/src/lexer"
	"github.com/MasaruFukazawa/monkey-lang/src/object"
	"github.com/MasaruFukazawa/monkey-lang/src/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {

	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluator := testEval(tt.input)
		testIntegerObject(t, evaluator, tt.expected)
	}

}

func testEval(input string) object.Object {

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {

	integer, ok := obj.(*object.Integer)

	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if integer.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", integer.Value, expected)
		return false
	}

	return true
}

func TestEvalBooleanExpression(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		evaluator := testEval(tt.input)
		testBooleanObject(t, evaluator, tt.expected)
	}

}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {

	result, ok := obj.(*object.Boolean)

	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}
