// As part of #loop exercise: map https://open.kattis.com/contests/xoh92u
// Easy 1.6 https://open.kattis.com/problems/conformity

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func conformity() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	combinations := map[string]int{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		combo := strings.Split(scanner.Text(), " ")
		sort.Strings(combo)
		comboStr := strings.Join(combo, " ")
		combinations[comboStr]++
	}

	count := 0
	highest := 0
	for _, val := range combinations {
		if val > highest {
			highest = val
			count = val
		} else if val == highest {
			count += val
		}
	}

	fmt.Println(count)
}
