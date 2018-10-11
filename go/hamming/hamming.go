package hamming

import "fmt"

// Distance counts hamming distance of input DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("disallow first strand longer, left(%d) != right(%d)",
			len(a), len(b))
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
