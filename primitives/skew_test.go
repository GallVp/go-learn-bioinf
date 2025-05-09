package primitives

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestComputeGCskew(t *testing.T) {
	got := computeGCskew("CATGGGCATCGGCCATACGCC")
	want := []int{0, -1, -1, -1, 0, 1, 2, 1, 1, 1, 0, 1, 2, 1, 0, 0, 0, 0, -1, 0, -1, -2}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
