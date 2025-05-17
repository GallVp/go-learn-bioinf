package primitives

import (
	"bytes"
	"io"
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

func TestSkewCmdMain(t *testing.T) {
	skewCmd := NewSkewCmd()

	// Capture output
	buf := bytes.NewBufferString("")
	skewCmd.SetOut(buf)
	skewCmd.SetErr(buf)
	skewCmd.SetArgs([]string{"-s", "CATGC"})

	skewCmd.Execute()
	out, _ := io.ReadAll(buf)

	got := string(out)
	want := "0\n-1\n-1\n-1\n0\n-1\n"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Mismatch (got - want):\n%s", diff)
	}
}
