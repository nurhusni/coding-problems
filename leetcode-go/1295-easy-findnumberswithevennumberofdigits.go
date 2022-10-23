package main

import "math"

func findNumbers(nums []int) int {
	count := 0

	for _, num := range nums {
		digits := math.Floor(math.Log10(float64(num))) + 1

		if int(digits)%2 == 0 {
			count++
		}
	}

	return count
}
