package main

import (
	"fmt"
	"math"
)

func main() {
	ssq := 0
	sqs := 0
	for i := 1; i < 101; i++ {
		ssq += int(math.Pow(float64(i), 2))
		sqs += i
	}

	fmt.Println(int(math.Pow(float64(sqs), 2)) - ssq)
}
