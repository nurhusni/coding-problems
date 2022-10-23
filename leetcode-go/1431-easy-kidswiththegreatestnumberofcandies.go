package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Kids With the Greatest Number of Candies.
// Memory Usage: 2.4 MB, less than 6.52% of Go online submissions for Kids With the Greatest Number of Candies.

func kidsWithCandies(candies []int, extraCandies int) []bool {
	mostCandies := 0

	for _, candy := range candies {
		if candy > mostCandies {
			mostCandies = candy
		}
	}

	result := []bool{}
	for _, candy := range candies {
		if candy+extraCandies >= mostCandies {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
