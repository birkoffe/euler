package main

import "fmt"

func main() {
	longestN := 0
	longest := 0

	cache := make(map[int]int)

	for x := 2; x < 1000001; x++ {
		n := x
		i := 0

		for {
			if n == 1 {
				break
			}

			k, ok := cache[n]
			if ok {
				i += k
				break
			}

			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}

			i++
		}

		cache[x] = i

		if i > longest {
			longest = i
			longestN = x
		}
	}

	fmt.Println(longestN)
}
