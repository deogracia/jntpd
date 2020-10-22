package cmd

import "testing"

func TestDumb(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}

}
