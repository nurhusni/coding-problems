package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mapNums := make(map[int]int)

	for i, num := range nums {
		if _, numExists := mapNums[num]; numExists {
			if i-mapNums[num] <= k {
				return true
			}
		}

		mapNums[num] = i
	}

	return false
}
