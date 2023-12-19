package flatarr

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindPimesInFib1(t *testing.T) {
	Convey("1", t, func() {
		arrs := make([]interface{}, 0)
		arrs = append(arrs, 9)
		arrs = append(arrs, []int{1, 2, 3})
		arrs = append(arrs, [][]int{{4, 5, 6}, {7}})
		arrs = append(arrs, 8)
		inner := make([]interface{}, 0)
		inner = append(inner, 10)
		inner = append(inner, []int{11, 12})
		outer := make([]interface{}, 0)
		outer = append(outer, inner...)
		arrs = append(arrs, outer)
		fmt.Println(arrs)
		res := FlaterArray(arrs)
		So(res, ShouldBeFalse)
	})
}
