package utils

func SplitByRunes(s string, n int) []string {
	r := []rune(s)
	var chunks []string
	for i := 0; i < len(r); i += n {
		end := min(i+n, len(r))
		chunks = append(chunks, string(r[i:end]))
	}
	return chunks
}
