package primes_fibonacci

import (
	"fmt"
	"math/big"
)

/*
	Given a number n, print all the prime numbers that are in the first n Fibonacci numbers.

	Input: n = 6
	Output: [2, 3, 5]

	Explanation: the first 6 fibonacci numbers are 1 1 2 3 5 8, and 2, 3, 5 are prime numbers
*/

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
