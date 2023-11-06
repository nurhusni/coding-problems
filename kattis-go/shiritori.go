// As part of #loop exercise: set https://open.kattis.com/contests/k6guuj
// Easy 2.5 https://open.kattis.com/problems/shiritori

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func shiritori() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	totalTurns, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	word := scanner.Text()
	lastLetter := string(word[len(word)-1])
	usedWords := map[string]struct{}{
		word: {},
	}

	for i := 2; i <= totalTurns; i++ {
		scanner.Scan()
		word := scanner.Text()

		if lastLetter != string(word[0]) {
			if i%2 == 0 {
				fmt.Println("Player 2 lost")
			} else {
				fmt.Println("Player 1 lost")
			}
			return
		}

		if _, exist := usedWords[word]; exist {
			if i%2 == 0 {
				fmt.Println("Player 2 lost")
			} else {
				fmt.Println("Player 1 lost")
			}
			return
		}

		lastLetter = string(word[len(word)-1])
		usedWords[word] = struct{}{}
	}

	fmt.Println("Fair Game")
}
