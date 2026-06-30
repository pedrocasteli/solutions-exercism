package differenceofsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	if n < 0 {
		return 0
	}

	sum := float64((n * (n + 1)) / 2)

	return int(math.Pow(sum, 2))
}

func SumOfSquares(n int) int {
	if n < 0 {
		return 0
	}

	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
