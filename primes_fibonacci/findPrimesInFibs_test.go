package primes_fibonacci

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindPimesInFib(t *testing.T) {
	Convey("1", t, func() {
		res := FindPrimesInFibonacciList(6)
		So(res, ShouldEqual, []int64{2, 3, 5})
	})
	Convey("2", t, func() {
		res := FindPrimesInFibonacciList(7)
		So(res, ShouldEqual, []int64{2, 3, 5, 13})
	})
	Convey("3", t, func() {
		res := FindPrimesInFibonacciList(8)
		So(res, ShouldEqual, []int64{2, 3, 5, 13})
	})
}
