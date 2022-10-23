package main

func productExceptSelf(nums []int) []int {
	leftProducts := make([]int, len(nums))
	rightProducts := make([]int, len(nums))

	for i, num := range nums {
		if i == 0 {
			leftProducts[0] = num
			continue
		}

		leftProducts[i] = leftProducts[i-1] * num
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			rightProducts[j] = nums[j]
			continue
		}

		rightProducts[j] = rightProducts[j+1] * nums[j]
	}

	res := []int{}
	for k := range nums {
		if k == 0 {
			res = append(res, rightProducts[k+1])
			continue
		}

		if k == len(nums)-1 {
			res = append(res, leftProducts[k-1])
			continue
		}

		res = append(res, leftProducts[k-1]*rightProducts[k+1])

	}

	return res
}
