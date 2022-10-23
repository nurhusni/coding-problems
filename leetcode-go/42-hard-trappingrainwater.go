package main

func trap(height []int) int {
	leftMax := []int{}
	rightMax := []int{}
	numberOfRainwater := 0

	for i := range height {
		// Populate the right max along with the left max
		rightMax = append(rightMax, 0)

		if i == 0 {
			leftMax = append(leftMax, 0)
			continue
		}

		if height[i-1] > leftMax[i-1] {
			leftMax = append(leftMax, height[i-1])
			continue
		}

		leftMax = append(leftMax, leftMax[i-1])
	}

	for j := len(height) - 1; j >= 0; j-- {
		if j == len(height)-1 {
			continue
		}

		if height[j+1] > rightMax[j+1] {
			rightMax[j] = height[j+1]
			continue
		}

		rightMax[j] = rightMax[j+1]
	}

	for k, val := range height {
		smallerHeight := 0
		if leftMax[k] <= rightMax[k] {
			smallerHeight = leftMax[k]
		} else {
			smallerHeight = rightMax[k]
		}

		rainwater := smallerHeight - val

		if rainwater > 0 {
			numberOfRainwater += rainwater
		}
	}

	return numberOfRainwater
}
