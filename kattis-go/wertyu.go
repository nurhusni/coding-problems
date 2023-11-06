// As part of #loop exercise: map https://open.kattis.com/contests/xoh92u
// Medium 3.3 https://open.kattis.com/problems/wertyu

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wertyu() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mapping := map[string]string{}
		characters := []string{"`1234567890-=", "qwertyuiop[]\\", "asdfghjkl;'", "zxcvbnm,./"}
		for _, word := range characters {
			for i, char := range word {
				if i == 0 {
					continue
				}

				mapping[string(char)] = string(word[i-1])
			}
		}

		text := scanner.Text()

		result := ""
		for _, letter := range text {
			if _, exist := mapping[strings.ToLower(string(letter))]; exist {
				result += mapping[strings.ToLower(string(letter))]
			} else {
				result += string(letter)
			}
		}

		fmt.Println(strings.ToUpper(result))
	}
}
