package main

func missingNumber(nums []int) int {
	numRange := make(map[int]int)

	for _, num := range nums {
		numRange[num] = 1
	}

	for i := 0; i <= len(nums); i++ {
		if numRange[i] == 0 {
			return i
		}
	}
	return 0
}
