package luhn

import "strings"	

func Valid(id string) bool {
	idNoSpaces := strings.ReplaceAll(id, " ", "")

	if len(idNoSpaces) <= 1 {
		return false
	}

	sum := 0
	length := len(idNoSpaces)

	for i, r := range idNoSpaces {
		if r < '0' || r > '9' {
			return false
		}

		digit := int(r - '0')

		if (length-1-i)%2 == 1 {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}
