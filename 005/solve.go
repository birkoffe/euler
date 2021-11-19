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

	max := 20
	result := 1

	for prime := range primes {
		k := 1
		if prime > max {
			break
		}
		for k*prime < max+1 {
			k *= prime
		}
		result *= k
	}

	fmt.Println(result)
}
