// Code generated by 'yaegi extract crypto/hmac'. DO NOT EDIT.

// +build go1.15

package stdlib

import (
	"crypto/hmac"
	"reflect"
)

func init() {
	Symbols["crypto/hmac"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Equal": reflect.ValueOf(hmac.Equal),
		"New":   reflect.ValueOf(hmac.New),
	}
}
