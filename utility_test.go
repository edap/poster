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
			So(isImage("image.jpg"), ShouldBeTrue)
			So(isImage("image.png"), ShouldBeTrue)
		})

		Convey("return true if the extension is an image extension, also for uppercase", func() {
			So(isImage("image.PNG"), ShouldBeTrue)
		})

		Convey("return false if the extension name is wrong written", func() {
			So(isImage("imagejpg"), ShouldBeFalse)
		})

		Convey("return false if the extension is not an image extension", func() {
			So(isImage("text.doc"), ShouldBeFalse)
		})
	})
}

func TestIsJpeg(t *testing.T) {
	Convey("Given a name file", t, func() {
		Convey("return true if the extension is a jpg extension", func() {
			So(isJpeg("image.jpg"), ShouldBeTrue)
		})

		Convey("return false if the extension is not a jpg", func() {
			So(isJpeg("text.doc"), ShouldBeFalse)
			So(isJpeg("image.png"), ShouldBeFalse)
		})
	})
}
