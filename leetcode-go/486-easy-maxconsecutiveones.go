package main

func findMaxConsecutiveOnes(nums []int) int {
	maxNum := 0
	count := 0

	for _, num := range nums {
		if num == 1 {
			count++

			if count > maxNum {
				maxNum = count
			}
		} else {
			count = 0
		}
	}

	return maxNum
}
