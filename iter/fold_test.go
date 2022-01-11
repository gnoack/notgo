package iter_test

import (
	"testing"

	"github.com/gnoack/notgo/iter"
)

func TestFold(t *testing.T) {
	sl := []string{"foo", "bar", "baz"}
	got := iter.Fold[string, string](iter.Slice[string](sl), "x", func(a, b string) string { return a + "," + b })
	want := "x,foo,bar,baz"

	if got != want {
		t.Errorf("iter.Fold(...) = %q, want %q", got, want)
	}
}
