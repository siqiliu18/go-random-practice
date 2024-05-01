package jumponcloud

/*
Input (stdin)
10
0 0 1 0 0 0 0 1 0 0

Expected Output
6
*/

func jumpingOnClouds(c []int32) int32 {
	// Write your code here     0 1 0 0 0 1 0
	arr := make([]int32, len(c))
	arr[0] = 0
	for ci := 1; ci < len(c); ci++ {
		if ci-2 < 0 {
			if c[ci] == 0 {
				arr[ci] = 1 // 0 0
			} else {
				arr[ci] = 0 // 0 1
			}
			continue
		}
		if c[ci-2] == 0 {
			if c[ci] == 0 { // 0 1 0
				arr[ci] = arr[ci-2] + 1
			} else { // 0 1 1
				arr[ci] = 0
			}
		}

		if c[ci-2] == 1 {
			arr[ci] = arr[ci-1] + 1
		}
	}

	return arr[len(arr)-1]
}
