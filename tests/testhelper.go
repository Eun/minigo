package tests

import (
	"bytes"
	"testing"

	"github.com/Eun/minigo/minigo"
)

func RunTest(t *testing.T, config minigo.Config, context interface{}, input string, expected string) {
	g, err := minigo.New(config)
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err = g.Run(bytes.NewReader([]byte(input)), context, &out); err != nil {
		t.Fatal(err)
	}
	if out.String() != expected {
		t.Fatalf("expected `%s' got `%s'", expected, out.String())
	}
}
