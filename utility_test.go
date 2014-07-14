package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRandStr(t *testing.T) {
	Convey("Given the lenght of the string to be generated", t, func() {
		Convey("generates a random string", func() {
			str := randStr(10)
			So(len(str), ShouldEqual, 10)
		})
	})
}

func TestIsImage(t *testing.T) {
	Convey("Given a name file", t, func() {

		Convey("return true if the extension is an image extension", func() {
			str := "image.jpg"
			So(isImage(str), ShouldBeTrue)
		})

		Convey("return true if the extension is an image extension", func() {
			str := "imagejpg"
			So(isImage(str), ShouldBeFalse)
		})

		Convey("return false if the extension is not image extension", func() {
			str := "text.doc"
			So(isImage(str), ShouldBeFalse)
		})
	})
}
