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

func TestAllApproxMatches(t *testing.T) {
	got := AllApproxMatches("GGGCCGTTGGT", "GGACCGTTGAC", 3)
	want := []int{0}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}

func TestAllApproxMatches2(t *testing.T) {
	got := AllApproxMatches("GGACCGTTGACGGACCGTTGAC", "GGACCGTTGAC", 0)
	want := []int{0, 11}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}

func TestAllApproxMatches3(t *testing.T) {
	got := AllApproxMatches("CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAAT", "ATTCTGGA", 3)
	want := []int{6, 7, 26, 27}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
