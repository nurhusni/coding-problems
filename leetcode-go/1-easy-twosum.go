package main

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}

	for i, num := range nums {
		secondValue := target - num
		if secondIndex, ok := hashMap[secondValue]; ok {
			return []int{i, secondIndex}
		}

		hashMap[num] = i
	}

	return []int{0, 0}
}
