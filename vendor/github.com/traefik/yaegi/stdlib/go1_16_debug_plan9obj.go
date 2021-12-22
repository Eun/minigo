// Code generated by 'yaegi extract debug/plan9obj'. DO NOT EDIT.

// +build go1.16,!go1.17

package stdlib

import (
	"debug/plan9obj"
	"go/constant"
	"go/token"
	"reflect"
)

func init() {
	Symbols["debug/plan9obj/plan9obj"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Magic386":   reflect.ValueOf(constant.MakeFromLiteral("491", token.INT, 0)),
		"Magic64":    reflect.ValueOf(constant.MakeFromLiteral("32768", token.INT, 0)),
		"MagicAMD64": reflect.ValueOf(constant.MakeFromLiteral("35479", token.INT, 0)),
		"MagicARM":   reflect.ValueOf(constant.MakeFromLiteral("1607", token.INT, 0)),
		"NewFile":    reflect.ValueOf(plan9obj.NewFile),
		"Open":       reflect.ValueOf(plan9obj.Open),

		// type definitions
		"File":          reflect.ValueOf((*plan9obj.File)(nil)),
		"FileHeader":    reflect.ValueOf((*plan9obj.FileHeader)(nil)),
		"Section":       reflect.ValueOf((*plan9obj.Section)(nil)),
		"SectionHeader": reflect.ValueOf((*plan9obj.SectionHeader)(nil)),
		"Sym":           reflect.ValueOf((*plan9obj.Sym)(nil)),
	}
}
