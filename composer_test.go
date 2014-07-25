package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCalcPrimeFactors(t *testing.T) {
	Convey("calculates the prime factors of a number", t, func() {
		Convey("when the number is > 0, return a slice containing the factors", func() {
			to_factorize := 20
			res := calcPrimeFactors(to_factorize)
			So(res[0], ShouldEqual, 2)
			So(res[1], ShouldEqual, 2)
			So(res[2], ShouldEqual, 5)
		})
		Convey("when the integer is null, the slice is empty", func() {
			slice := []int{}
			to_factorize := 0
			res := calcPrimeFactors(to_factorize)
			So(res, ShouldHaveSameTypeAs, slice)
		})
	})
}

func TestCalculatePositions(t *testing.T) {
	Convey("Given an array containing the path of the images, the dimension of the rectangle and desired width and height return an array that contains two numbers", t, func() {
		imgs := []string{"test", "test1", "test2", "test4"}
		rect := map[string]int{
			"base":   2,
			"height": 2,
		}
		positions := calculatePositions(rect, imgs, 100, 100)
		So(positions["test"][0], ShouldEqual, 0)
		So(positions["test4"][1], ShouldEqual, -100)
	})
}

func TestGetBaseAndHeight(t *testing.T) {
	Convey("Given an array of prime factors, return an array that contains two numbers", t, func() {
		r := []int{2, 3}
		found, b, h := getBaseAndHeight(r)
		So(found, ShouldBeTrue)
		So(b, ShouldEqual, 3)
		So(h, ShouldEqual, 2)

	})
}

func TestIsASquare(t *testing.T) {
	Convey("Given a square number, return true and the side of the square", t, func() {
		sqr, side := isASquare(9)
		So(sqr, ShouldBeTrue)
		So(side, ShouldEqual, 3)
	})

	Convey("Given a non square number, return false and the side of the square", t, func() {
		sqr, side := isASquare(10)
		So(sqr, ShouldBeFalse)
		So(side, ShouldEqual, 3.1622776601683795)
	})
}

func TestcalculateRectangle(t *testing.T) {
	Convey("Given a number representing the area of a rectang, calculate base and height", t, func() {
		Convey("when it is the pow of a number, return base and height", func() {
			res := map[string]int{
				"area":    16,
				"height":  0,
				"base":    0,
				"skipped": 0,
			}
			rect := calculateRectangle(res)
			So(rect["height"], ShouldEqual, 4)
			So(rect["base"], ShouldEqual, 4)
			So(rect["skipped"], ShouldEqual, 0)
		})
		Convey("when it is not a prime number return base and height", func() {
			res := map[string]int{
				"area":    6,
				"height":  0,
				"base":    0,
				"skipped": 0,
			}
			rect := calculateRectangle(res)
			So(rect["height"], ShouldEqual, 2)
			So(rect["base"], ShouldEqual, 3)
			So(rect["skipped"], ShouldEqual, 0)
		})
		Convey("when is a prime number, remove one integer until is possible to find base and height", func() {
			res := map[string]int{
				"area":    19,
				"height":  0,
				"base":    0,
				"skipped": 0,
			}
			rect := calculateRectangle(res)
			So(rect["height"], ShouldEqual, 2)
			So(rect["base"], ShouldEqual, 9)
			So(rect["skipped"], ShouldEqual, 1)
		})
	})
}
