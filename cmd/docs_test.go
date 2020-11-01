package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestDocsAction(t *testing.T) {

	tempDir := t.TempDir()
	out := bytes.NewBufferString("")

	err := docsAction(out, tempDir)
	if err != nil {
		t.Fatalf("Expected \"%s\"; got \"%s\". tempDir = \"%s\"", "nil", out.String(), tempDir)
	}

	_, err = os.Stat(tempDir)
	if os.IsNotExist(err) {
		t.Fatalf("\"%s\" doesn't exist.", tempDir)
	}

	matches, err := filepath.Glob(tempDir + "*/*.md")
	if err != nil {
		t.Fatal("Something went wrong with filepath.Glob")
	}

	if len(matches) <= 0 {
		t.Fatalf("Expected at least one mardown file generated. Got \"%d\".", len(matches))
	}

}
