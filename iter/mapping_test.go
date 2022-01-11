package iter_test

import (
	"strconv"
	"testing"

	"github.com/gnoack/notgo/iter"
	"github.com/google/go-cmp/cmp"
)

func TestMapList(t *testing.T) {
	got := iter.ToSlice[string](iter.Map[int, string](iter.Slice[int]([]int{1, 2, 42}), strconv.Itoa))
	want := []string{"1", "2", "42"}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Unexpected map result, diff:\n%v", diff)
	}
}
