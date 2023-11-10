package primes_fibonacci

import (
	"fmt"
	"math/big"
)

func FindPrimesInFibonacciList(n int64) []int64 {
	dp := make([]int64, n)
	dp[0] = 1
	dp[1] = 1

	res := new([]int64)

	currFib(n-1, dp, res)

	fmt.Println(dp)

	return *res
}

func currFib(n int64, dp []int64, primes *[]int64) {
	if n < 2 {
		return
	}

	var i int64 = 2
	for ; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		if big.NewInt(dp[i]).ProbablyPrime(0) {
			*primes = append(*primes, dp[i])
		}
	}
}
