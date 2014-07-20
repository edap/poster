package main

import (
	_ "errors"
	//"fmt"
	_ "github.com/nfnt/resize"
	. "github.com/smartystreets/goconvey/convey"
	//"image"
	_ "image/color"
	//"image/jpeg"
	//"os"
	"testing"
)

// func TestForceToJpg(t *testing.T) {
// 	Convey("When a given image is not")
// }

func TestHasDesiredDimension(t *testing.T) {
	Convey("Check if an image has the same desired dimension", t, func() {
		//default_thumb, _ := openThumb("test_images/120x90.jpg")

		Convey("if the given dimensions are different from the image dimension, return false", func() {
			thumb := &Thumb{
				width:          120,
				height:         90,
				desired_width:  10,
				desired_height: 20,
			}
			result, _ := thumb.HasDesiredDimension()
			So(result, ShouldBeFalse)
		})

		Convey("if the given dimensions are the same as the image dimension, return true", func() {
			thumb := &Thumb{
				width:          120,
				height:         90,
				desired_width:  120,
				desired_height: 90,
			}
			result, _ := thumb.HasDesiredDimension()
			So(result, ShouldBeTrue)
		})

	})
}

func TestCopy(t *testing.T) {
	Convey("Check if copy file between folders works", t, func() {
		thumb := &Thumb{}
		_, err := thumb.Copy("test_images/120x90.jpg", "test_images/120x90copy.jpg")
		So(err, ShouldBeNil)
	})
}

func TestMove(t *testing.T) {
	Convey("Check if move files works", t, func() {
		thumb := &Thumb{}
		err := thumb.Move("test_images/120x90copy.jpg", "test_images/120x90moved.jpg")
		So(err, ShouldBeNil)
	})
}

func TestScale(t *testing.T) {
	Convey("Check if an image has the same desired dimension", t, func() {
		Convey("if the given dimensions are the same as the image dimension, return true", func() {
			_, err := openThumb("test_images/120x90.jpg")
			if err != nil {
				panic("impossible to open test image")
			}
			thumb := &Thumb{
				width:          120,
				height:         90,
				desired_width:  220,
				desired_height: 190,
			}
			er := thumb.Scale("test_images/120x90.jpg")

			So(er, ShouldBeNil)
		})
	})
}

// Convey("When one of the given param is < than 1", t, func() {
//  Convey("return a wrongArgumentError", func() {
//    _, err := HasDesiredDimension(120, 0, 120, 90)
//    w := new(wrongArgumentError)
//    So(err, ShouldHaveSameTypeAs, w)
//  })
//  Convey("give a significant message", func() {
//    _, err := HasDesiredDimension(120, 0, 120, 90)
//    So(err.Error(), ShouldEqual, "The argument thumb_height can not be minor than 1")
//  })
// })

// vedere qui http://golangtutorials.blogspot.de/2011/06/memory-variables-in-memory-and-pointers.html
// per l'errore cannot take the address of "test_image"
// Convey("if the given dimensions are the same as the image dimension, return true", func() {
//  source_dir = &"test_image"
//  target_dir = &"merged/test_image"
//  thumb := &Thumb{
//    width:          120,
//    height:         90,
//    desired_width:  300,
//    desired_height: 100,
//  }
//  result, _ := createImage("test_image/a_big_one.jpg")
//  thumb.ScaleThumb()
//  So(result, ShouldBeTrue)
// })

// func TestScaleThumb(t *testing.T) {

// }
