package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringBins(t *testing.T) {
	got := StringBins("CATGGGCATCGGCCATACGCC", 4)
	want := []string{"CATG", "GGCA", "TCGG", "CCAT", "ACGC", "C"}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}

func TestStringSlidingSlices(t *testing.T) {
	got := StringSlidingSlices("CATGG", 2)
	want := []string{"CA", "AT", "TG", "GG"}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}

func TestStringSlidingSlices2(t *testing.T) {
	got := StringSlidingSlices("C", 1)
	want := []string{"C"}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}

func TestStringSlidingSlices3(t *testing.T) {
	got := StringSlidingSlices("C", 0)
	want := []string{}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}

func TestStringSlidingSlices4(t *testing.T) {
	got := StringSlidingSlices("", 2)
	want := []string{}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
