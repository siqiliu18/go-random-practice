package main

import (
	"fmt"
	rm "go-random-practice/rotate_matrix"
	"sort"
)

func main() {
	fmt.Println("ready")

	inputMatrix := [][]int64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(rm.RotateMatrixRight90(inputMatrix))

	arr := []int{3, 3, 4, 5, 2, 2, 3, 2}
	m := make(map[int]int)
	for _, num := range arr {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num] += 1
		}
	}
	keys := []int{}
	for key := range m {
		keys = append(keys, key)
	}
	// sort.SliceStable will sort the keys of the map by default
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	fmt.Println(keys)
}
