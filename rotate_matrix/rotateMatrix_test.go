package rotate_matrix

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotateMatrixRight90(t *testing.T) {
	Convey("1", t, func() {
		inputMatrix := [][]int64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		res := RotateMatrixRight90(inputMatrix)
		So(len(res), ShouldEqual, 3)
		So(res[0][0], ShouldEqual, 7)
	})
	Convey("2", t, func() {
		inputMatrix := [][]int64{
			{1, 2, 3, 10},
			{4, 5, 6, 11},
			{7, 8, 9, 12},
		}
		res := RotateMatrixRight90(inputMatrix)
		So(len(res), ShouldEqual, 4)
		So(res[3][0], ShouldEqual, 12)
		So(res[3][2], ShouldEqual, 10)
	})
}

func TestRotateMatrixLeft90(t *testing.T) {
	Convey("1", t, func() {
		inputMatrix := [][]int64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		res := RotateMatrixLeft90(inputMatrix)
		fmt.Println(res)
		So(len(res), ShouldEqual, 3)
		So(res[0][0], ShouldEqual, 3)
	})
	Convey("2", t, func() {
		inputMatrix := [][]int64{
			{1, 2, 3, 10},
			{4, 5, 6, 11},
			{7, 8, 9, 12},
		}
		res := RotateMatrixLeft90(inputMatrix)
		So(len(res), ShouldEqual, 4)
		So(res[3][0], ShouldEqual, 1)
		So(res[3][2], ShouldEqual, 7)
	})
}
