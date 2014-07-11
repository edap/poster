package main

import (
	"fmt"
	"github.com/nfnt/resize"
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

type Image interface {
	HasDesiredDimension() (bool, error)
	CurrentWidth() int
	CurrentHeight() int
	Move(src_path, dst_path string) error
	Copy(src_path, dst_path string) (int64, error)
	Scale(img_path string) error
}

type Thumb struct {
	width          int
	height         int
	desired_width  int
	desired_height int
	img_name       string
}

func NewThumb(img_width int, img_height int, thumb_width int, thumb_height int, img_name string) Image {
	return &Thumb{
		width:          img_width,
		height:         img_height,
		desired_width:  thumb_width,
		desired_height: thumb_height,
		img_name:       img_name,
	}
}

func (t *Thumb) CurrentWidth() int {
	return t.width
}

func (t *Thumb) CurrentHeight() int {
	return t.height
}

func (t *Thumb) Scale(img_path string) error {
	are_equal, err := t.HasDesiredDimension()
	if err != nil {
		return err
	}
	if are_equal == true {
		return nil
	}

	file, err := os.Open(img_path)
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(uint(t.desired_width), uint(t.desired_height), img, resize.NearestNeighbor)
	out, err := os.Create(img_path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	return jpeg.Encode(out, m, nil)
}

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

func (t *Thumb) HasDesiredDimension() (bool, error) {
	if t.desired_height < 1 {
		return false, &wrongArgumentError{"thumb_height"}
	}
	if t.desired_width < 1 {
		return false, &wrongArgumentError{"thumb_width"}
	}

	if t.desired_width == t.width && t.desired_height == t.height {
		return true, nil
	} else {
		return false, nil
	}
}
