package primitives

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindIndices(t *testing.T) {
	skew := computeGCskew("TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGAT")
	got := findIndicesOfMin(skew)
	want := []int{11, 24}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
