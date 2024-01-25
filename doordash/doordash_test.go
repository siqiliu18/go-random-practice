package doordash

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeFormats(t *testing.T) {
	Convey("1", t, func() {
		res := TimeFormats("mon 10:00 am", "mon 2:00 pm")
		So(res, ShouldEqual, "")
	})
}
