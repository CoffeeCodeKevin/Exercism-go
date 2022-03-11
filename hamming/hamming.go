package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hammingCount := 0
	if len(a) != len(b) {
		return 0, errors.New("mismatched strand lengths")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingCount++
		}
	}
	return hammingCount, nil
}
