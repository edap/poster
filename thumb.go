package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	//"path/filepath"
	//errors
)

// questo sparira' non ha senso dire che 0 e' un error, quando di default, il pacchetto resize diche che se
// h o w sono 0, scala l'immagine in proporzione
type wrongArgumentError struct{ arg string }

func (a *wrongArgumentError) Error() string {
	return fmt.Sprintf("The argument %s can not be minor than 1", a.arg)
}

type Img interface {
	HasDesiredDimension() bool
	Height() int
	Width() int
	SetHeight(h int)
	SetWidth(w int)
	SetDimension() error
	GetFormatFromExtension() (error, string)
	Move(src_path, dst_path string) error
	Copy(src_path, dst_path string) (int64, error)
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
	// implementare gestione degli errori
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

func (t *Thumb) Copy(src_path, dst_path string) (int64, error) {
	src_file, err := os.Open(src_path)
	if err != nil {
		return 0, err
	}
	defer src_file.Close()

	src_file_stat, err := src_file.Stat()
	if err != nil {
		return 0, err
	}
	if !src_file_stat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src_path)
	}

	dst_file, err := os.Create(dst_path)
	if err != nil {
		return 0, err
	}
	defer dst_file.Close()
	return io.Copy(dst_file, src_file)
}

func (t *Thumb) Move(src_path, dst_path string) error {
	return os.Rename(src_path, dst_path)
}

func (t *Thumb) HasDesiredDimension() bool {
	if t.desired_width == t.width && t.desired_height == t.height {
		return true
	} else {
		return false
	}
}
