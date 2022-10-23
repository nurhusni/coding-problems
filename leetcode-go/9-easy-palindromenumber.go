package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	xStr := fmt.Sprint(x)
	iter := len(xStr)

	for i, digit := range xStr {
		if i == len(xStr)/2 {
			return true
		}

		iter--

		if byte(digit) != xStr[iter] {
			return false
		}
	}

	return true
}
