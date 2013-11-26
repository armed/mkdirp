package mkdirp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var c = Convey

func TestMkdirp(t *testing.T) {
	c("mk func should", t, func() {
		c("return list of paths to create", func() {
			result := mk("/Users/me/somefolder")

			So(len(result), ShouldEqual, 1)
			So(result[0], ShouldEqual, "/Users/me/somefolder")
		})
		c("support creating of trees", func() {
			result := mk("/Users/me/{somefolder,somefolder2}")

			So(len(result), ShouldEqual, 2)
			So(result[0], ShouldEqual, "/Users/me/somefolder")
			So(result[1], ShouldEqual, "/Users/me/somefolder2")
		})
	})
}
