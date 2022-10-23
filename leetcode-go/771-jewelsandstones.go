package main

func numJewelsInStones(jewels string, stones string) int {
	mapJewels := map[string]bool{}

	for _, jewel := range jewels {
		mapJewels[string(jewel)] = true
	}

	count := 0

	for _, stone := range stones {
		if _, stoneExists := mapJewels[string(stone)]; stoneExists {
			count++
		}
	}

	return count
}
