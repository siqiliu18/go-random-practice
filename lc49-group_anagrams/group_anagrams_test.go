package lc49groupanagrams

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindPimesInFib1(t *testing.T) {
	Convey("1", t, func() {
		arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		res := groupAnagrams(arr)
		exp := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
		So(res, ShouldEqual, exp)
	})
}

func TestFindPimesInFib2(t *testing.T) {
	Convey("1", t, func() {
		arr := []string{"eat", "tea", "tan"}
		res := groupAnagrams(arr)
		exp := [][]string{{"tan"}, {"eat", "tea"}}
		So(res, ShouldEqual, exp)
	})
}

func TestFindPimesInFib3(t *testing.T) {
	Convey("1", t, func() {
		arr := []string{"", ""}
		res := groupAnagrams(arr)
		exp := [][]string{{"", ""}}
		So(res, ShouldEqual, exp)
	})
}

func TestFindPimesInFib4(t *testing.T) {
	Convey("1", t, func() {
		arr := []string{"ab", "ab"}
		res := groupAnagrams(arr)
		exp := [][]string{{"ab", "ab"}}
		So(res, ShouldEqual, exp)
	})
}

func TestFindPimesInFib5(t *testing.T) {
	Convey("1", t, func() {
		arr := []string{"", "b"}
		res := groupAnagrams(arr)
		exp := [][]string{{""}, {"b"}}
		So(res, ShouldEqual, exp)
	})
}

func TestFindPimesInFib6(t *testing.T) {
	Convey("1", t, func() {
		arr := []string{"tea", "and", "ace", "ad", "eat", "dans"}
		res := groupAnagrams(arr)
		exp := [][]string{{"dans"}, {"ad"}, {"ace"}, {"and"}, {"eat", "tea"}}
		So(res, ShouldEqual, exp)
	})
}
