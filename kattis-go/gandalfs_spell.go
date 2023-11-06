// As part of #loop exercise: map https://open.kattis.com/contests/xoh92u
// Easy 2.2 https://open.kattis.com/problems/gandalfsspell

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func gandalfsSpell() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	keys := scanner.Text()
	mapLetterSentence := map[string]string{}
	mapSentenceLetter := map[string]string{}

	scanner.Scan()
	sentenceString := strings.Split(scanner.Text(), " ")

	if len(sentenceString) != len(keys) {
		fmt.Println("False")
		return
	}

	for i, word := range sentenceString {
		if _, exist := mapSentenceLetter[word]; exist {
			if string(keys[i]) != mapSentenceLetter[word] {
				fmt.Println("False")
				return
			}
			continue
		} else {
			mapSentenceLetter[word] = string(keys[i])
			if _, exist := mapLetterSentence[string(keys[i])]; exist {
				if mapLetterSentence[string(keys[i])] != word {
					fmt.Println("False")
					return
				}
			}
			mapLetterSentence[string(keys[i])] = word
		}
	}

	fmt.Println("True")
}
