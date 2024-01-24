package doordash

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeFormats(t *testing.T) {
	Convey("1", t, func() {
		res := TimeFormats()
		So(res, ShouldEqual, "")
	})
}
