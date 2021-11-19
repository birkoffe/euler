package main

import "fmt"

func polindrome(i int) int {
	result := 0

	for i > 0 {
		result = result*10 + i%10
		i /= 10
	}

	return result
}

func main() {
	result := 0

	for i := 10; i < 1000; i++ {
		for j := 10; j < 1000; j++ {
			ij := i * j
			if ij == polindrome(ij) && ij > result {
				result = ij
			}
		}
	}

	fmt.Println(result)
}
