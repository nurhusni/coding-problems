package main

func findDisappearedNumbers(nums []int) []int {
	appear := make(map[int]int)
	disappear := []int{}

	for _, num := range nums {
		appear[num] = 1
	}

	for i := 1; i <= len(nums); i++ {
		if appear[i] == 0 {
			disappear = append(disappear, i)
		}
	}

	return disappear
}
