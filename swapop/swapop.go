package swapop

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// swap once between any indexes
func maximumSwap(num int) int {
	arr := []int{}
	for ; num > 0; num /= 10 {
		curr := num % 10
		arr = append(arr, curr)
	}
	// for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
	// 	arr[i], arr[j] = arr[j], arr[i]
	// }
	fmt.Println(arr)
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			arr[i], arr[j] = arr[j], arr[i]
			num := arrToNum(arr)
			if num > max {
				max = num
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return max
}

func arrToNum(arr []int) int {
	num := 0
	i := 0
	n := 0
	for i < len(arr) {
		num = num + arr[i]*(int(math.Pow10(n)))
		i++
		n++
	}
	return num
}

// Given an array a[ ] and the number of adjacent swap operations allowed are K. The task is to find the max number that can be formed using these swap operations.
/*
	1. Find the max val in arr, move it close to the front using k times.
	2. If there is remain rounds, find the next max val.

	repeat these 2 steps till k turns to 0
*/

func maximizeArr(arr []int, k int) []int {
	if k == 0 {
		return nil
	}

	for i := 0; i < len(arr); i++ {
		maxi := 0
		max := math.MinInt
		limit := 0
		if k+i > len(arr) {
			limit = len(arr)
		} else {
			limit = k + i
		}

		for j := i; j <= limit; j++ {
			if arr[j] > max {
				max = arr[j]
				maxi = j
			}
		}

		k = k - (maxi - i)

		swapMax(arr, i, maxi)

		if k == 0 {
			return arr
		}
	}
	return arr
}

func swapMax(arr []int, target, curr int) {
	aux := 0
	for i := curr; i > target; i-- {
		aux = arr[i-1]
		arr[i-1] = arr[i]
		arr[i] = aux
	}
}

func getMinNumMoves(blocks []int32) int32 {
	// Write your code here
	toSort := []int32{}

	// insert blocks into toSort array
	for _, val := range blocks {
		toSort = append(toSort, val)
	}

	// sort array for the max and min val
	sort.SliceStable(toSort, func(i, j int) bool {
		return toSort[i] < toSort[j]
	})

	arrLen := len(toSort)
	maxv := toSort[arrLen-1]
	minv := toSort[0]

	mini, maxi := 0, 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == minv {
			mini = i
		}
		if blocks[i] == maxv {
			maxi = i
		}
	}

	res := mini - 1

	if maxi < mini {
		res = res + arrLen - maxi - 1
	} else {
		res = res + arrLen - maxi
	}

	return int32(res)
}
/* 
	- another source: https://leetcode.com/discuss/interview-experience/803109/amazon-phone-interview-sde2-aug-2020

	E.G. s = "asdfoo123ate" and forbbiden ="[foo, ate]". Answer is 7 ("oo123at").
	    findReviewScore("FastDeliveryOkayProduct", []string{"eli", "yo", "eryok"}) => 11 (ayProduct)
*/
func findReviewScore(review string, prohibitedWords []string) int32 {
	// lowercase
	lowerReview := strings.ToLower(review)

	proIndexMap := make(map[int]int)
	for _, s := range prohibitedWords {
		i := strings.Index(lowerReview, s) 
		proIndexMap[i] = i + len(s)
	}
	fmt.Println(proIndexMap)

	res := 0
	begin := 0
	curser := 0
	for ; curser < len(review); {
		if endIndex, ok := proIndexMap[curser]; ok { // found index
			for innerIndex := curser + 1; innerIndex < endIndex; innerIndex++ {
				if innerWordEndIndex, ok := proIndexMap[innerIndex]; ok { // found nested word
					currLen := innerWordEndIndex - curser - 1
					if currLen > res {
						res = currLen
					}
					begin = innerIndex + 1
					curser = begin
					break
				}
			}
			currLen := endIndex - begin - 1
			if currLen > res {
				res = currLen
			}
			begin = curser + 1
			curser = begin
		} else {
			curser++
		}
	}
	fmt.Printf("begin = %v\n", begin)
	if begin < len(review) {
		currLen := len(review) - begin
		if currLen > res {
			res = currLen
		}
	}
	return int32(res)
}
