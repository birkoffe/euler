package main

import "fmt"

func main() {
	max := 1000

	for a := 1; a < max/3; a++ {
		for b := a + 1; b < (max-a)/2; b++ {
			c := max - a - b

			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
