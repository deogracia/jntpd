package cmd

import (
	"bytes"
	"os"
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

}
