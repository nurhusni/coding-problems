// As part of #loop exercise: map https://open.kattis.com/contests/xoh92u
// Medium 4.1 https://open.kattis.com/problems/addingwords

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addingWords() {
	scanner := bufio.NewScanner(os.Stdin)
	wordnum := map[string]int{}

	for scanner.Scan() {
		if scanner.Text() == "clear" {
			wordnum = map[string]int{}
			continue
		}

		command := strings.Split(scanner.Text(), " ")
		if command[0] == "def" {
			num, _ := strconv.Atoi(command[2])

			wordnum[command[1]] = num
		} else if command[0] == "calc" {
			calculation := 0
			result := strings.Join(command[1:], " ")
			if _, exist := wordnum[command[1]]; exist {
				calculation = wordnum[command[1]]
				for i := 2; i < len(command); i++ {
					if command[i] == "=" {
						resultString := "unknown"
						for key, val := range wordnum {
							if calculation == val {
								resultString = key
								break
							}
						}
						result += fmt.Sprintf(" %s", resultString)
						break
					}

					if command[i] == "+" || command[i] == "-" || command[i] == " " {
						continue
					}

					if _, exist := wordnum[command[i]]; !exist {
						result += " unknown"
						break
					}

					if command[i-1] == "+" {
						calculation += wordnum[command[i]]
					} else if command[i-1] == "-" {
						calculation -= wordnum[command[i]]
					}
				}
			} else {
				result += " unknown"
			}

			fmt.Println(result)
		}
	}
}
