package numbers

import "math"

func FindPrimeNumbers(n float64) []float64 {
	numbers := []float64{}

	isPrime := true
	if n == 2 {
		numbers = append(numbers, 2)
	}

	for i := 2.0; i < n; i++ {
		for _, number := range numbers {
			dev := i / number
			if dev == math.Trunc(dev) {
				isPrime = false
				break
			}
		}
		if isPrime {
			numbers = append(numbers, i)
		} else {
			isPrime = true
			continue
		}
	}
	return numbers
}
