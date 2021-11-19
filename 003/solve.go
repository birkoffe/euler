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
	number := 600851475143
	max_prime := 0

	primes := make(chan int)

	go func() {
		for i := 2; ; i++ {
			if isPrime(i) {
				primes <- i
			}
		}
	}()

	for number != 1 {
		prime := <-primes
		if number%prime == 0 {
			number /= prime
			max_prime = prime
		}
	}

	fmt.Println(max_prime)
}
