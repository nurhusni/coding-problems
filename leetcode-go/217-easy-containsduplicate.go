package main

func containsDuplicate(nums []int) bool {
	mapNums := make(map[int]bool)

	for _, num := range nums {
		if _, numExists := mapNums[num]; numExists {
			return true
		} else {
			mapNums[num] = true
		}
	}

	return false
}
