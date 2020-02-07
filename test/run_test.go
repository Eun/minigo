package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// build minigo
	err := exec.Command("go", "build", "-o=minigo", "github.com/Eun/minigo").Run()
	if err != nil {
		t.Fatal("go build failed", err)
		return
	}

	var testFiles []string
	// collect all go files and run them
	err = filepath.Walk("./tests/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".minigo") {
			return nil
		}
		testFiles = append(testFiles, path)
		return nil
	})

	for _, testFile := range testFiles {
		t.Run(testFile, func(t *testing.T) {
			out, err := exec.Command("./minigo", testFile).CombinedOutput()
			if err != nil {
				t.Fatalf("%v: output is %s", err, string(out))
			}
			if string(out) != "Hello World\n" {
				t.Fatalf("expected %s got %s", "Hello World\n", string(out))
			}
		})
	}

	if err != nil {
		t.Fatal(err)
	}
}
