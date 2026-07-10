package isbnverifier

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	cleanIsbn := strings.ReplaceAll(isbn, "-", "")

	if len(cleanIsbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		if cleanIsbn[i] < '0' || cleanIsbn[i] > '9' {
			return false
		}
		digit := int(cleanIsbn[i] - '0')
		sum += digit * (10 - i)
	}

	last := cleanIsbn[9]
	var lastValue int

	switch {
	case last == 'X':
		lastValue = 10
	case last >= '0' && last <= '9':
		lastValue = int(last - '0')
	default:
		return false
	}

	sum += lastValue

	return sum%11 == 0
}
