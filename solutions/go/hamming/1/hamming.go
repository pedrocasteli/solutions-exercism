package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	strandA := strings.TrimSpace(strings.ToUpper(a))
	strandB := strings.TrimSpace(strings.ToUpper(b))
	differences := 0

	if len(strandA) == len(strandB) {
		for i := 0; i < len(strandA); i++ {
			if strandA[i] != strandB[i] {
				differences++
			}
		}

		return differences, nil
	} else {
		return differences, errors.New("The sequences have different lengths")
	}
}
