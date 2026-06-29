package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be a positive integer")
	} else {
		if n == 1 {
			return 0, nil
		}

		steps := 0
		result := n

		for {
			if result%2 == 0 {
				result = result / 2
				steps++
			} else {
				result = result * 3
				result += 1
				steps++
			}

			if result == 1 {
				break
			}
		}

		return steps, nil
	}
}
