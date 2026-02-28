package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Different length")
	}
	var errorCount int
	for i := range len(a) {
		if b[i] != a[i] {
			errorCount++
		}
	}
	return errorCount, nil
}
