package cmd

import (
	"bytes"
	"testing"
)

func TestDocsAction(t *testing.T) {

	tempDir := t.TempDir()
	out := bytes.NewBufferString("")

	err := docsAction(out, "/tmp/jntpdn-test")
	if err != nil {
		t.Fatalf("Expected \"%s\"; got \"%s\". tempDir = \"%s\"", "nil", out.String(), tempDir)
	}
}
