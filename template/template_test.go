package template

import (
	"testing"
)

func TestNew(t *testing.T) {
	templates, err := New(nil)
	if err != nil {
		t.Errorf("unexpected error creating template: %s", err.Error())
		return
	}
	if templates == nil {
		t.Error("expected templates, got: nil")
		return

	}
}
