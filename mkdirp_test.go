package mkdirp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var c = Convey

func TestMkdirp(t *testing.T) {
	c("mkTree should", t, func() {
		c("return tree from straight path", func() {
			root := mkTree("Users/me/somefolder")
			So(root, ShouldNotBeNil)
			So(len(root.branches), ShouldEqual, 0)
			So(root.name, ShouldEqual, "Users/me/somefolder")
		})
		c("return tree with subtrees", func() {
			root := mkTree("Users/me/{somefolder,somefolder2}/test/{data1,data2}")

			So(root.name, ShouldEqual, "Users/me/")
			subtree := root.branches
			So(len(subtree), ShouldEqual, 2)
			So(subtree[0].name, ShouldEqual, "somefolder")
			So(subtree[1].name, ShouldEqual, "somefolder2")

			sstree := subtree[0]
			sstree2 := subtree[1]
			So(len(sstree.branches), ShouldEqual, 1)
			So(sstree.branches[0].name, ShouldEqual, "test/")
			So(len(sstree2.branches), ShouldEqual, 1)
			So(sstree2.branches[0].name, ShouldEqual, "test/")

			test := sstree.branches[0]
			So(len(test.branches), ShouldEqual, 2)
			So(test.branches[0].name, ShouldEqual, "data1")
			So(test.branches[1].name, ShouldEqual, "data2")

		})
	})
	// c("mk func should", t, func() {
	// 	c("return list of paths to create", func() {
	// 		result := mk("/Users/me/somefolder")

	// 		So(len(result), ShouldEqual, 1)
	// 		So(result[0], ShouldEqual, "/Users/me/somefolder")
	// 	})
	// 	c("support creating of trees", func() {
	// 		result := mk("tmpdir/{trunk/sources/{includes,docs},branches,tags}")
	// 		fmt.Println(result)
	// 		So(len(result), ShouldEqual, 4)
	// 		So(result[0], ShouldEqual, "tmpdir/trunk/sources/includes")
	// 		So(result[1], ShouldEqual, "tmpdir/trunk/sources/docs")
	// 		So(result[2], ShouldEqual, "tmpdir/trunk/branches")
	// 		So(result[3], ShouldEqual, "tmpdir/trunk/tags")
	// 	})
	// })
}
