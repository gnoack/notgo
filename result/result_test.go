package result_test

import (
	"io"
	"testing"

	"github.com/gnoack/notgo/result"
)

func TestResultOk(t *testing.T) {
	r := result.Of("hello")
	if r.Err() != nil {
		t.Errorf("result.Of(...) may not be an error, but got: %v", r.Err())
	}
	want := "hello"
	if r.Value() != want {
		t.Errorf("r.Value() = %q, want %q", r.Value(), want)
	}
}

func TestResultError(t *testing.T) {
	r := result.Error[string](io.EOF)
	if r.Err() != io.EOF {
		t.Errorf("r.Err() = %q, want %q", r.Err(), io.EOF)
	}
	if r.Value() != "" {
		t.Errorf("r.Value() = %q, want %q", r.Value(), "")
	}
}
