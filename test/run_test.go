package test

import (
	"fmt"
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

		out, err := exec.Command("./minigo", path).CombinedOutput()
		if err != nil {
			return fmt.Errorf("test %s failed: %v\noutput is %s", path, err, string(out))
		}
		if string(out) != "Hello World\n" {
			t.Fatalf("test %s failed, expected %s got %s", path, "Hello World\n", string(out))
			return fmt.Errorf("test %s failed", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
