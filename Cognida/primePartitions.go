package main

import (
	"fmt"
	"math"
)

func sieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func primePartitions(n int) int {
	const MOD = 1000000007
	primes := sieveOfEratosthenes(n)

	dp := make([]int, n+1)
	dp[0] = 1

	for _, prime := range primes {
		for i := prime; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-prime]) % MOD
		}
	}

	return dp[n]
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(primePartitions(N))
}
