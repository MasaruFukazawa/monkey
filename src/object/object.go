/**
 * パッケージ名: object
 * ファイル名: object.go
 * 概要: オブジェクトを定義する
 */
package object

type ObjectType string

// オブジェクトの種類を定義する
type Object interface {
	Type() ObjectType
	Inspect() string
}
