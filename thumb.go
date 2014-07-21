package main

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
)

type Img interface {
	HasDesiredDimension() bool
	Height() int
	Width() int
	SetHeight(h int)
	SetWidth(w int)
	SetDimension() error
	GetFormatFromExtension() (error, string)
	DecodeIt() (image.Image, error)
}

type Thumb struct {
	width          int
	height         int
	desired_width  int
	desired_height int
	img_name       string
}

//func NewThumb(img_width int, img_height int, thumb_width int, thumb_height int, img_name string) Img {
func NewThumb(thumb_width int, thumb_height int, img_name string) Img {
	return &Thumb{
		desired_width:  thumb_width,
		desired_height: thumb_height,
		img_name:       img_name,
	}
}

func (t *Thumb) SetDimension() error {
	img_file, err := os.Open(t.img_name)
	defer img_file.Close()
	if err != nil {
		log.Printf("the image can not be opened: %v", err)
		return err
	}
	config, _, err := image.DecodeConfig(img_file)
	if err != nil {
		log.Printf("the image %s can not be decoded: %v", t.img_name, err)
		return err
	}
	t.SetHeight(config.Height)
	t.SetWidth(config.Width)

	return err
}

func (t *Thumb) DecodeIt() (image.Image, error) {
	img_file, err := os.Open(t.img_name)
	defer img_file.Close()
	if err != nil {
		log.Printf("the image can not be opened: %v", err)
		return nil, err
	}

	img, _, err := image.Decode(img_file)
	if err != nil {
		log.Printf("that image can not be decoded: %v", err)
		return nil, err
	}

	if t.HasDesiredDimension() {
		return img, nil
	} else {
		m := resize.Resize(uint(t.desired_width), uint(t.desired_height), img, resize.NearestNeighbor)
		return m, nil
	}
}

func (t *Thumb) GetFormatFromExtension() (error, string) {
	return nil, "jpeg"
}

func (t *Thumb) Width() int {
	return t.width
}

func (t *Thumb) Height() int {
	return t.height
}

func (t *Thumb) SetWidth(width int) {
	t.width = width
}

func (t *Thumb) SetHeight(height int) {
	t.height = height
}

func (t *Thumb) forceToJpg(w io.Writer, r io.Reader) error {
	// custom error, the image can not be converted to jpg
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, nil)
}

// decodificare sempre in jpg
// convertToPNG converts from any recognized format to PNG.
// func convertToPNG(w io.Writer, r io.Reader) error {
//  img, _, err := image.Decode(r)
//  if err != nil {
//   return err
//  }
//  return png.Encode(w, img)
// }

func (t *Thumb) HasDesiredDimension() bool {
	if t.desired_width == t.width && t.desired_height == t.height {
		return true
	} else {
		return false
	}
}
