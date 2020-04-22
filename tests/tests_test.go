package tests

import (
	"testing"

	"github.com/Eun/minigo/minigo"
)

func TestSimple(t *testing.T) {
	runTest(t,
		minigo.Config{
			TemplateMode: false,
		},
		struct{ Name string }{"Joe"},
		`fmt.Printf("Hello %s", context.Name)`,
		"Hello Joe",
	)
}

func TestAdvanced(t *testing.T) {
	runTest(t,
		minigo.Config{
			TemplateMode: false,
		},
		struct{ Name string }{"Joe"},
		`package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello", name)
}

func main() {
	hello(context.Name)
}
`,
		"Hello Joe\n",
	)
}

func TestImport(t *testing.T) {
	runTest(t,
		minigo.Config{
			TemplateMode: false,
		},
		struct{ Name string }{"Joe"},
		`package main

import "./world"

func main() {
	fmt.Printf("Hello %s's %s\n", context.Name, world.World())
}
`,
		"Hello Joe's World\n",
	)
}

func TestTemplate(t *testing.T) {
	runTest(t,
		minigo.Config{
			TemplateMode: true,
		},
		struct{ Name string }{"Joe"},
		`Hello <$ fmt.Println(context.Name) $>`,
		"Hello Joe\n",
	)
}

func TestShebangSkip(t *testing.T) {
	runTest(t,
		minigo.Config{
			TemplateMode: false,
		},
		struct{ Name string }{"Joe"},
		`#!panic("THIS SHOULD BE IGNORED")
fmt.Println("Hello", context.Name)
`,
		"Hello Joe\n",
	)
}
