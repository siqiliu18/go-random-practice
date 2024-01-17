package lc56mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	il, ir := 0, 1

	endi := len(intervals)

	lv := intervals[il][0]
	rv := intervals[il][1]

	for il < endi {
		if ir >= endi {
			curr := []int{lv, rv}
			res = append(res, curr)
			return res
		}

		if intervals[ir][0] <= rv && rv <= intervals[ir][1] {
			rv = intervals[ir][1]
			ir++
		} else if rv >= intervals[ir][1] {
			ir++
		} else {
			curr := []int{lv, rv}
			res = append(res, curr)
			il = ir
			ir = il + 1
			lv = intervals[il][0]
			rv = intervals[il][1]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
