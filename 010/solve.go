package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	primes := make(chan int)

	go func() {
		for i := 2; ; i++ {
			if isPrime(i) {
				primes <- i
			}
		}
	}()

	prime := <-primes
	sum := 0

	for prime < 2000000 {
		sum += prime
		prime = <-primes
	}

	fmt.Println(sum)
}
