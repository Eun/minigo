// Code generated by 'goexports unicode/utf8'. DO NOT EDIT.

// +build go1.13,!go1.14

package stdlib

import (
	"reflect"
	"unicode/utf8"
)

func init() {
	Symbols["unicode/utf8"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DecodeLastRune":         reflect.ValueOf(utf8.DecodeLastRune),
		"DecodeLastRuneInString": reflect.ValueOf(utf8.DecodeLastRuneInString),
		"DecodeRune":             reflect.ValueOf(utf8.DecodeRune),
		"DecodeRuneInString":     reflect.ValueOf(utf8.DecodeRuneInString),
		"EncodeRune":             reflect.ValueOf(utf8.EncodeRune),
		"FullRune":               reflect.ValueOf(utf8.FullRune),
		"FullRuneInString":       reflect.ValueOf(utf8.FullRuneInString),
		"MaxRune":                reflect.ValueOf(utf8.MaxRune),
		"RuneCount":              reflect.ValueOf(utf8.RuneCount),
		"RuneCountInString":      reflect.ValueOf(utf8.RuneCountInString),
		"RuneError":              reflect.ValueOf(utf8.RuneError),
		"RuneLen":                reflect.ValueOf(utf8.RuneLen),
		"RuneSelf":               reflect.ValueOf(utf8.RuneSelf),
		"RuneStart":              reflect.ValueOf(utf8.RuneStart),
		"UTFMax":                 reflect.ValueOf(utf8.UTFMax),
		"Valid":                  reflect.ValueOf(utf8.Valid),
		"ValidRune":              reflect.ValueOf(utf8.ValidRune),
		"ValidString":            reflect.ValueOf(utf8.ValidString),
	}
}
