package main

import (
	_ "errors"
	"fmt"
	"image"
	_ "image/color"
	"image/jpeg"
	"os"
)

func createImage(name string) (bool, error) {
	m := image.NewRGBA(image.Rect(0, 0, 120, 90))

	out, err := os.Create(name)
	if err != nil {
		panic(fmt.Sprintf("is not possible to create the file %s necessary for testing", name))
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
	return true, err
}

func openThumb(img_path string) (*os.File, error) {
	if file, err := os.Open(img_path); err == nil {
		defer file.Close()
		return file, nil
	} else {
		panic(fmt.Sprintf("is not possible to open the file %s necessary for testing", img_path))
	}
}