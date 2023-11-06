// As part of #loop exercise: set https://open.kattis.com/contests/k6guuj
// Medium 4.8 https://open.kattis.com/problems/cd

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cd() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		nm := scanner.Text()

		var n, m int
		nmslice := strings.Split(nm, " ")
		n, _ = strconv.Atoi(nmslice[0])
		if len(nmslice) > 1 {
			m, _ = strconv.Atoi(nmslice[1])
		}

		if n == 0 && m == 0 {
			return
		}

		values := map[int]struct{}{}
		for i := 0; i < n; i++ {
			scanner.Scan()
			val, _ := strconv.Atoi(scanner.Text())

			values[val] = struct{}{}
		}

		count := 0
		for j := 0; j < m; j++ {
			scanner.Scan()
			val, _ := strconv.Atoi(scanner.Text())

			if _, exist := values[val]; exist {
				count++
			}
		}

		fmt.Println(count)
	}
}
