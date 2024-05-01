package nychaos

import (
	"fmt"
	// "math"
)

/*
https://www.hackerrank.com/challenges/new-year-chaos/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

It is New Year's Day and people are in line for the Wonderland rollercoaster ride. Each person wears a sticker indicating their initial position in the queue from  to . Any person can bribe the person directly in front of them to swap positions, but they still wear their original sticker. One person can bribe at most two others.

Determine the minimum number of bribes that took place to get to a given queue order. Print the number of bribes, or, if anyone has bribed more than two people, print Too chaotic.

*/

func minimumBribes(q []int32) int32 {
	// Write your code here
	var res int32 = 0
	for i, num := range q {
		diff := num - int32(i) - 1
		if diff > 2 {
			fmt.Println("Too chaotic")
			return -1
		}
		for j := maxInt(0, int(q[i])-2); j < i; j++ {
			if q[j] > q[i] {
				res++
			}
		}
	}
	return res
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
