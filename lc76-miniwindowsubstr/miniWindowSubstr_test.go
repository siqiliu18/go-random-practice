package lc76miniwindowsubstr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinWindow1(t *testing.T) {
	Convey("1", t, func() {
		s := "ADOBECODEBANC"
		t := "ABC"
		res := minWindowAns(s, t)
		exp := "BANC"
		So(res, ShouldEqual, exp)
	})
}

func TestMinWindow2(t *testing.T) {
	Convey("1", t, func() {
		s := "a"
		t := "a"
		res := minWindowAns(s, t)
		exp := "a"
		So(res, ShouldEqual, exp)
	})
}

func TestMinWindow3(t *testing.T) {
	Convey("1", t, func() {
		s := "a"
		t := "aa"
		res := minWindowAns(s, t)
		exp := ""
		So(res, ShouldEqual, exp)
	})
}

func TestMinWindow4(t *testing.T) {
	Convey("1", t, func() {
		s := "bba"
		t := "ab"
		res := minWindowAns(s, t)
		exp := "ba"
		So(res, ShouldEqual, exp)
	})
}

func TestMinWindow5(t *testing.T) {
	Convey("1", t, func() {
		s := "bbaa"
		t := "aba"
		res := minWindowAns(s, t)
		exp := "baa"
		So(res, ShouldEqual, exp)
	})
}
