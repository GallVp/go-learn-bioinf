package primitives

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNumberOfMismatches(t *testing.T) {
	got, _ := NumberOfMismatches("GGGCCGTTGGT", "GGACCGTTGAC")
	want := 3

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
