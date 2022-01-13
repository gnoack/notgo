package slices_test

import (
	"strings"
	"testing"

	"github.com/gnoack/notgo/slices"
	"github.com/google/go-cmp/cmp"
)

func TestSelect(t *testing.T) {
	for _, tt := range []struct {
		Input     []string
		Want      []string
		Predicate func(string) bool
	}{
		{
			Input:     []string{"foo", "bar", "baz", "quux"},
			Predicate: func(s string) bool { return strings.Contains(s, "a") },
			Want:      []string{"bar", "baz"},
		},
		{
			Input:     []string{"foo", "bar", "baz", "quux"},
			Predicate: func(s string) bool { return len(s) < 4 },
			Want:      []string{"foo", "bar", "bazz"},
		},
	} {
		got := slices.Select(tt.Input, tt.Predicate)
		if diff := cmp.Diff(got, tt.Want); diff != "" {
			t.Errorf("slices.Select(%q, predicate) = %q, want %q", tt.Input, got, tt.Want)
		}
	}
}
