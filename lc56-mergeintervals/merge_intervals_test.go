package lc56mergeintervals

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeIntervals1(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		res := merge(input)
		exp := [][]int{{1, 6}, {8, 10}, {15, 18}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals2(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 7}, {2, 6}, {8, 17}, {15, 18}, {16, 20}, {22, 25}, {23, 27}}
		res := merge(input)
		exp := [][]int{{1, 7}, {8, 20}, {22, 27}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals3(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 7}, {2, 6}, {5, 17}, {18, 20}, {22, 25}, {23, 27}}
		res := merge(input)
		exp := [][]int{{1, 17}, {18, 20}, {22, 27}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals4(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 30}, {2, 6}, {5, 17}, {18, 20}, {22, 25}, {23, 27}}
		res := merge(input)
		exp := [][]int{{1, 30}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals5(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 7}, {3, 6}, {2, 17}, {18, 20}, {22, 25}, {23, 27}}
		res := merge(input)
		exp := [][]int{{1, 17}, {18, 20}, {22, 27}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals6(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 7}, {1, 6}}
		res := merge(input)
		exp := [][]int{{1, 7}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals7(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 3}}
		res := merge(input)
		exp := [][]int{{1, 3}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals8(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 3}}
		res := merge(input)
		exp := [][]int{{1, 3}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals9(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}
		res := merge(input)
		exp := [][]int{{2, 4}, {5, 5}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals10(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 4}, {0, 2}, {3, 5}}
		res := merge(input)
		exp := [][]int{{0, 5}}
		So(res, ShouldEqual, exp)
	})
}

func TestMergeIntervals11(t *testing.T) {
	Convey("1", t, func() {
		input := [][]int{{1, 4}, {1, 4}}
		res := merge(input)
		exp := [][]int{{1, 4}}
		So(res, ShouldEqual, exp)
	})
}
