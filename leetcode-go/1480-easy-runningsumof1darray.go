package main

func runningSum(nums []int) []int {
	count := 0
	result := []int{}

	for _, num := range nums {
		count += num
		result = append(result, count)
	}

	return result
}
