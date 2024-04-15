package possiblesum

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPossiblesum1(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{10, 50, 100}
		quanity := []int{1, 2, 1}
		res := solution(coins, quanity)
		exp := 9
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum2(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{10, 50, 100, 500}
		quanity := []int{5, 3, 2, 2}
		res := solution(coins, quanity)
		exp := 122
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum3(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1}
		quanity := []int{5}
		res := solution(coins, quanity)
		exp := 5
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum4(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1, 1}
		quanity := []int{2, 3}
		res := solution(coins, quanity)
		exp := 5
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum5(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1, 2}
		quanity := []int{50000, 2}
		res := solution(coins, quanity)
		exp := 50004
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum6(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1, 2, 3}
		quanity := []int{2, 3, 10000}
		res := solution(coins, quanity)
		exp := 30008
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum7(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{3, 1, 1}
		quanity := []int{111, 84, 104}
		res := solution(coins, quanity)
		exp := 521
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum21(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{10, 50, 100}
		quanity := []int{1, 2, 1}
		res := solution2(coins, quanity)
		exp := 9
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum22(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{10, 50, 100, 500}
		quanity := []int{5, 3, 2, 2}
		res := solution2(coins, quanity)
		exp := 122
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum23(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1}
		quanity := []int{5}
		res := solution2(coins, quanity)
		exp := 5
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum24(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1, 1}
		quanity := []int{2, 3}
		res := solution2(coins, quanity)
		exp := 5
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum25(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1, 2}
		quanity := []int{50000, 2}
		res := solution2(coins, quanity)
		exp := 50004
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum26(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{1, 2, 3}
		quanity := []int{2, 3, 10000}
		res := solution2(coins, quanity)
		exp := 30008
		So(res, ShouldEqual, exp)
	})
}

func TestPossiblesum27(t *testing.T) {
	Convey("1", t, func() {
		coins := []int{3, 1, 1}
		quanity := []int{111, 84, 104}
		res := solution2(coins, quanity)
		exp := 521
		So(res, ShouldEqual, exp)
	})
}
