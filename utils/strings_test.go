package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplitByRunes(t *testing.T) {
	got := SplitByRunes("CATGGGCATCGGCCATACGCC", 4)
	want := []string{"CATG", "GGCA", "TCGG", "CCAT", "ACGC", "C"}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
