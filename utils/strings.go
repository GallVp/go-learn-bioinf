package utils

import "github.com/thoas/go-funk"

// Bin a string into 'x' slices of length 'n' except the last one
// which contains the remaining runes
func StringBins(s string, n int) []string {
	r := []rune(s)
	var chunks []string
	for i := 0; i < len(r); i += n {
		end := min(i+n, len(r))
		chunks = append(chunks, string(r[i:end]))
	}
	return chunks
}

// Return sliding slices of length 'l' for a string
func StringSlidingSlices(s string, l int) []string {

	if len(s) < 1 {
		return []string{}
	}

	if l < 1 {
		return []string{}
	}

	r := []rune(s)
	var slices []string
	for i := 0; i < len(r)-l+1; i += 1 {
		slices = append(slices, string(r[i:i+l]))
	}
	return slices
}

func StringsWithIndices(strings []string) []funk.Tuple {
	indices := make([]int, len(strings))
	for i := range strings {
		indices[i] = i
	}

	return funk.Zip(strings, indices)
}
