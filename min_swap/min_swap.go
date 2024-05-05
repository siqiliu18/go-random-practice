package minswap

/*
Question:
You are given an unordered array consisting of consecutive integers  [1, 2, 3, ..., n] without any duplicates. You are allowed to swap any two elements. Find the minimum number of swaps required to sort the array in ascending order.

Source:
https://www.hackerrank.com/challenges/minimum-swaps-2/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
*/

/*
The minimumSwaps function takes in an array of integers arr and returns the minimum number of swaps required to sort the array. The function works by creating a map m that maps each element in the array to its index in the array. Then, it iterates over the entire array and swaps any elements that are out of place. The swaps are tracked using the variable s, which keeps track of the total number of swaps made. Finally, the function returns the value of s.

In this example, we assume that the input array is [4, 3, 1, 2]. The first iteration of the loop will check if the current element is equal to its expected position (which is always true for the first element). Since the element is not in its expected position, the second iteration of the loop will swap the element with the smallest value that is greater than the current element. In this case, the smallest value greater than 4 is 5, so the element will be swapped with 5. The updated array becomes [5, 3, 1, 2]. The map m will also be updated to reflect the new positions of the elements.

The next iteration of the loop will check if the current element is still in its expected position. This time, the element is not in its expected position, so the third iteration of the loop will swap the element with the smallest value that is greater than the current element. In this case, the smallest value greater than 5 is 6, so the element will be swapped with 6. The updated array becomes [6, 3, 1, 2]. The map m will also be updated to reflect the new positions of the elements.

The process continues until all elements have been sorted. In this example, the minimum number of swaps required would be 2, as two swaps were made: [4, 3, 1, 2] -> [3, 4, 1, 2] -> [3, 1, 4, 2].
*/

func minimumSwaps(arr []int32) int32 {
	n := len(arr)
	var s int32 = 0
	m := make(map[int32]int32)
	for i := 0; i < n; i++ {
		m[arr[i]] = int32(i)
	}
	for i := 0; i < n; i++ {
		min := int32(i + 1)
		if min != arr[i] {
			minIndex := m[min]
			m[arr[i]] = minIndex
			arr[minIndex] = arr[i]
			s += 1
		}
	}
	return s
}
