package main

import (
	_ "errors"
	"fmt"
	_ "github.com/nfnt/resize"
	. "github.com/smartystreets/goconvey/convey"
	//"image"
	_ "image/color"
	"image/jpeg"
	"os"
	"testing"
)

func TestHasDesiredDimension(t *testing.T) {
	Convey("Check if an image has the same desired dimension", t, func() {

		Convey("if the given dimensions are different from the image dimension, return false", func() {
			thumb := &Thumb{
				width:          120,
				height:         90,
				desired_width:  10,
				desired_height: 20,
			}
			result := thumb.HasDesiredDimension()
			So(result, ShouldBeFalse)
		})

		Convey("if the given dimensions are the same as the image dimension, return true", func() {
			thumb := &Thumb{
				width:          120,
				height:         90,
				desired_width:  120,
				desired_height: 90,
			}
			result := thumb.HasDesiredDimension()
			So(result, ShouldBeTrue)
		})

	})
}

func TestDecodeIt(t *testing.T) {
	Convey("Open an image and decode it", t, func() {
		original_img := "test_images/120x90.jpg"
		final_img := "test_images/decode.jpg"
		Convey("if the desired dimensions are different as the original one, scale it", func() {
			thumb := &Thumb{
				img_name:       original_img,
				desired_width:  20,
				desired_height: 10,
			}
			img, er := thumb.DecodeIt()
			So(er, ShouldBeNil)

			out, err := os.Create(final_img)
			if err != nil {
				panic(fmt.Sprintf("is not possible to create the file %s necessary for testing", final_img))
			}
			defer out.Close()
			jpeg.Encode(out, img, nil)

			final_file, err := os.Open(final_img)
			if err != nil {
				fmt.Println(err)
			}
			config, _ := jpeg.DecodeConfig(final_file)
			So(config.Width, ShouldEqual, 20)
			So(config.Height, ShouldEqual, 10)

		})
	})
}

// func TestSetDimension(t *testing.T){
// 	Convey("read the dimension of an image and set width and height"){

// 	}

// }
