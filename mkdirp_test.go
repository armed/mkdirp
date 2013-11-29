package mkdirp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var c = Convey

func TestMkdirp(t *testing.T) {
	c("mkTree should", t, func() {
		c("return tree from straight path", func() {
			root := mkTree("/Users/me/somefolder")
			So(root, ShouldNotBeNil)
			So(len(root.getChildren()), ShouldEqual, 0)
			So(root.getName(), ShouldEqual, "/Users/me/somefolder")
			So(len(root.getPaths()), ShouldEqual, 1)
			So(root.getPaths()[0], ShouldEqual, "/Users/me/somefolder")
		})
		c("create directory which name contains one charachter", func() {
			root := mkTree("a")
			So(root.getName(), ShouldEqual, "a")
			So(len(root.getPaths()), ShouldEqual, 1)
			So(root.getPaths()[0], ShouldEqual, "a")
		})
		c("return tree with subtrees", func() {
			root := mkTree("/Users/me/{somefolder,somefolder2}/test/{data1,data2}")

			So(root.getName(), ShouldEqual, "/Users/me/")
			subtree := root.getChildren()
			So(len(subtree), ShouldEqual, 2)
			So(subtree[0].getName(), ShouldEqual, "somefolder")
			So(subtree[1].getName(), ShouldEqual, "somefolder2")

			sstree := subtree[0]
			sstree2 := subtree[1]
			So(len(sstree.getChildren()), ShouldEqual, 1)
			So(sstree.getChildren()[0].getName(), ShouldEqual, "test/")
			So(len(sstree2.getChildren()), ShouldEqual, 1)
			So(sstree2.getChildren()[0].getName(), ShouldEqual, "test/")

			test := sstree.getChildren()[0]
			So(len(test.getChildren()), ShouldEqual, 2)
			So(test.getChildren()[0].getName(), ShouldEqual, "data1")
			So(test.getChildren()[1].getName(), ShouldEqual, "data2")

			c("where root contains slice of all paths", func() {
				paths := root.getPaths()
				So(len(paths), ShouldEqual, 4)
				So(paths[0], ShouldEqual, "/Users/me/somefolder/test/data1")
				So(paths[1], ShouldEqual, "/Users/me/somefolder/test/data2")
				So(paths[2], ShouldEqual, "/Users/me/somefolder2/test/data1")
				So(paths[3], ShouldEqual, "/Users/me/somefolder2/test/data2")
			})
		})
	})
}
