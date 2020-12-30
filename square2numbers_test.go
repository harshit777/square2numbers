package square2numbers_test

import (
	"testing"

	"github.com/harshit777/square2numbers"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExecute(t *testing.T) {
	Excecute := square2numbers.Excecute
	Convey("Should convert correctly", t, func() {
		Convey("Small Number should convert correctly",
			func() {
				So(Excecute(1), ShouldEqual, 1)
				So(Excecute(2), ShouldEqual, 4)
				So(Excecute(3), ShouldEqual, 9)
			})
	})

}
