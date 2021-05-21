package cmd

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocsCmd(t *testing.T) {
	wanted := "*cobra.Command"
	cmd := docsCmd()

	if assert.NotNil(t, cmd) {
		if got := reflect.TypeOf(cmd).String(); got != wanted {
			log.Fatalf("Expected \"%s\"; but got \"%s\" variable", wanted, got)
		}
	}
}
