/**
 * パッケージ名: object
 * ファイル名: object.go
 * 概要: オブジェクトを定義する
 */
package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
)

// オブジェクトの種類を定義する
type Object interface {
	Type() ObjectType
	Inspect() string
}

// 整数オブジェクトを表す構造体
type Integer struct {
	Value int64
}

// 整数オブジェクトの種類を返す
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// 整数オブジェクトの値を返す
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// 真偽値オブジェクトを表す構造体
type Boolean struct {
	Value bool
}

// 真偽値オブジェクトの種類を返す
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// 真偽値オブジェクトの値を返す
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
