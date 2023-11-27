package cuttingribbons

// This logic works but slow, see the 3rd test case - leetcode says timeout
func MaxLength(ribbons []int, k int) int {
	maxVal := 1
	for _, val := range ribbons {
		if val > maxVal {
			maxVal = val
		}
	}

	maxLens := []int{}
	for i := 1; i <= maxVal; i++ {
		currSum := 0
		for _, num := range ribbons {
			if i > num {
				continue
			}
			val := num / i
			currSum += val
		}
		if currSum < k && len(maxLens) > 0 {
			return maxLens[len(maxLens)-1]
		} else if currSum >= k {
			maxLens = append(maxLens, i)
		}
	}
	if len(maxLens) > 0 {
		return maxLens[len(maxLens)-1]
	}
	return 0
}

// https://github.com/doocs/leetcode/blob/main/solution/1800-1899/1891.Cutting%20Ribbons/README_EN.md
func MaxLength2(ribbons []int, k int) int {
	left := 0
	right := 0
	for _, val := range ribbons {
		if val > right {
			right = val
		}
	}
	for left < right {
		mid := (left + right + 1) >> 1
		cnt := 0
		for _, x := range ribbons {
			if mid > x {
				continue
			}
			cnt += x / mid
		}
		if cnt >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
