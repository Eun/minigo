// Code generated by 'yaegi extract crypto/sha1'. DO NOT EDIT.

// +build go1.16

package stdlib

import (
	"crypto/sha1"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["crypto/sha1/sha1"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BlockSize": reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"New":       reflect.ValueOf(sha1.New),
		"Size":      reflect.ValueOf(constant.MakeFromLiteral("20", token.INT, 0)),
		"Sum":       reflect.ValueOf(sha1.Sum),
	}
}
