package swapop

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaximumSwap(t *testing.T) {
	Convey("1", t, func() {
		res := maximumSwap(2736)
		So(res, ShouldEqual, 7236)
	})
	Convey("1", t, func() {
		res := maximumSwap(9973)
		So(res, ShouldEqual, 9973)
	})
}

func TestArrToNum(t *testing.T) {
	Convey("1", t, func() {
		res := arrToNum([]int{6, 3, 7, 2})
		So(res, ShouldEqual, 2736)
	})
}

func TestMaximizeArr(t *testing.T) {
	Convey("1", t, func() {
		res := maximizeArr([]int{1, 2, 4, 3}, 2)
		So(res, ShouldEqual, []int{4, 1, 2, 3})
	})
}

func TestAWS(t *testing.T) {
	Convey("0", t, func() {
		res := getMinNumMoves([]int32{2, 4, 3, 1, 6})
		So(res, ShouldEqual, 3)
	})
}

func TestFindReviewScore1(t *testing.T) {
	Convey("0", t, func() {
		res := findReviewScore("FastDeliveryOkayProduct", []string{"eli", "yo", "eryok"})
		So(res, ShouldEqual, 10)
	})
}

func TestFindReviewScore2(t *testing.T) {
	Convey("0", t, func() {
		res := findReviewScore("asdfoo123ate", []string{"foo", "ate"})
		So(res, ShouldEqual, 7)
	})
}
