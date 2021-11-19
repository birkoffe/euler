package main

import "fmt"

func main() {
	sum, a, b := 0, 1, 2

	for a < 4000000 {
		if a%2 == 0 {
			sum += a
		}

		a, b = b, a+b
	}

	fmt.Println(sum)
}
