package main

import (
	"fmt"
)

func factorsCount(triangle int) int {
	count := 0

	for i := 1; i < triangle+1; i++ {
		if triangle%i == 0 {
			count++
		}
	}

	return count
}

func slow() {
	// 200s in M1
	triangle := 0
	n := 0

	for {
		n += 1
		triangle += n

		count := factorsCount(triangle)

		if count >= 500 {
			fmt.Println(triangle, count)
			return
		}
	}
}

func fast() {
	// TODO: write fast algo
}

func main() {
	// slow()

	fast()
}
