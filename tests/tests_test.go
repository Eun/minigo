package tests

import (
	"testing"

	"github.com/Eun/minigo/pkg/minigo"
)

func TestSimple(t *testing.T) {
	runTest(t,
		minigo.Config{
			StartTokens: []rune{},
			EndTokens:   []rune{},
		},
		struct{ Name string }{"Joe"},
		`print("Hello " + context.Name)`,
		"Hello Joe",
	)
}

func TestAdvanced(t *testing.T) {
	runTest(t,
		minigo.Config{
			StartTokens: []rune{},
			EndTokens:   []rune{},
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
			StartTokens: []rune{},
			EndTokens:   []rune{},
		},
		struct{ Name string }{"Joe"},
		`package main

import "fmt"
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
			StartTokens: []rune("<$"),
			EndTokens:   []rune("$>"),
		},
		struct{ Name string }{"Joe"},
		`Hello <$ println(context.Name) $>`,
		"Hello Joe\n",
	)
}

func TestShebangSkip(t *testing.T) {
	runTest(t,
		minigo.Config{
			StartTokens: []rune{},
			EndTokens:   []rune{},
		},
		struct{ Name string }{"Joe"},
		`#!panic("THIS SHOULD BE IGNORED")
println("Hello " + context.Name)
`,
		"Hello Joe\n",
	)
}
