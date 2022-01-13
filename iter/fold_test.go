package iter_test

import (
	"testing"

	"github.com/gnoack/notgo/iter"
)

func TestFold(t *testing.T) {
	sl := []string{"foo", "bar", "baz"}
	concat := func(a, b string) string { return a + "," + b }
	got := iter.Fold(iter.Slice(sl), "x", concat)
	want := "x,foo,bar,baz"

	if got != want {
		t.Errorf("iter.Fold(...) = %q, want %q", got, want)
	}
}
