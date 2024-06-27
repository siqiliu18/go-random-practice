package lc253meetingrooms

// https://leetcode.com/problems/meeting-rooms-ii/

import (
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	start := []int{}
	end := []int{}

	for _, interval := range intervals {
		start = append(start, interval[0])
		end = append(end, interval[1])
	}

	sort.SliceStable(start, func(i, j int) bool {
		return start[i] < start[j]
	})

	sort.SliceStable(end, func(i, j int) bool {
		return end[i] < end[j]
	})

	si, ei := 0, 0

	count := 0
	for si < len(start) {
		if start[si] < end[ei] {
			count++
			si++
		} else {
			count--
			ei++
		}
	}

	return count
}

/* from WCA
type Interval struct {
	start, end int
}

type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	n := len(intervals)
	intvs := make([]Interval, n)
	for i := 0; i < n; i++ {
		intvs[i] = Interval{intervals[i][0], intervals[i][1]}
	}

	sort.Sort(ByStart(intvs))

	rooms := 0
	end := intvs[0].end
	for i := 1; i < n; i++ {
		if intvs[i].start >= end {
			rooms++
			end = intvs[i].end
		}
	}

	return rooms
}
*/
