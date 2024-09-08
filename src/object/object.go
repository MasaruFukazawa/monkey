/**
 * パッケージ名: object
 * ファイル名: object.go
 * 概要: オブジェクトを定義する
 */
package object

//import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
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
