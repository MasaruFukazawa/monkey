/**
 * パッケージ名: object
 * ファイル名: object.go
 * 概要: オブジェクトを定義する
 */
package object

import (
	"bytes"
	"fmt"
	"github.com/MasaruFukazawa/monkey-lang/src/ast"
	"hash/fnv"
	"strings"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

// オブジェクトの種類を定義する
type Object interface {
	Type() ObjectType
	Inspect() string
}

// オブジェクトがハッシュキーとして使えることを示すインターフェース
type Hashable interface {
	HashKey() HashKey
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

// NULLオブジェクトを表す構造体
type Null struct{}

// NULLオブジェクトの種類を返す
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// NULLオブジェクトの値を返す
func (n *Null) Inspect() string {
	return "null"
}

// 戻り値オブジェクトを表す構造体
type ReturnValue struct {
	Value Object
}

// 戻り値オブジェクトの種類を返す
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// 戻り値オブジェクトの値を返す
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// エラーオブジェクトを表す構造体
type Error struct {
	Message string
}

// エラーオブジェクトの種類を返す
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// エラーオブジェクトの値を返す
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// 関数オブジェクトを表す構造体
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// 関数オブジェクトの種類を返す
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// 関数オブジェクトの値を返す
func (f *Function) Inspect() string {

	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {")
	out.WriteString(f.Body.String())
	out.WriteString("¥n}")

	return out.String()
}

// 文字列オブジェクトを表す構造体
type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

// 組み込み関数オブジェクトを表す構造体
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}

// 配列オブジェクトを表す構造体
type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (ao *Array) Inspect() string {

	var out bytes.Buffer

	elements := []string{}

	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// ハッシュキーオブジェクトを表す構造体
type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {

	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {

	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (h *Hash) Inspect() string {

	var out bytes.Buffer

	pairs := []string{}

	for _, pair := range h.Pairs {

		pairs = append(pairs, fmt.Sprintf(
			"%s: %s",
			pair.Key.Inspect(),
			pair.Value.Inspect(),
		))

	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
