package main

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for i := range accounts {
		totalWealth := 0

		for _, wealth := range accounts[i] {
			totalWealth += wealth
		}

		if totalWealth > maxWealth {
			maxWealth = totalWealth
		}
	}

	return maxWealth
}
