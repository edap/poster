package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

//Gomerger is a command line tool that allows you to merge more images in one.
//Given a source folder containing the images, the program resizes all the images at the
//same dimension (default is 120x90) and calculates the disposition of the images in a rectangle.
//If the number of the images is a prime number, and one of the image can not fit into a rectangle,
//the tool will skip it, until the total number of the images will fit into the rectangle.
func main() {
	var (
		thumb_width  = flag.Int("thumb_width", 120, "the width of a single thumb")
		thumb_height = flag.Int("thumb_height", 90, "the height of a single thumb")
		source_dir   = flag.String("source_dir", ".", "the origin directory that contains the images to compose the grid")
		dest_dir     = flag.String("dest_dir", ".", "the destination directory that will contain the grid")
		log_file     = flag.String("log_file", "stdout", "specify a log file, as default it will print on stdout")
	)
	flag.Parse()

	// set a log file if it's required
	if *log_file != "stdout" {
		f, err := os.OpenFile(*log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	// At least 2 images has to be in the folder present
	tot, images := listFiles(*source_dir)
	if tot < 2 {
		log.Fatal("There are less than two images in this folder")
	}

	// create the destination directory
	destination := createDirectory(*dest_dir)
	if destination != nil {
		log.Fatalf("impossible to create destination directory: %v", destination)
	}

	//calculate the dimension of the rectangle
	res := map[string]int{"area": tot, "height": 0, "base": 0, "skipped": 0}
	rect := calculateRectangle(res)
	log.Printf("%d images will be skipped", res["skipped"])
	log.Printf("%d images will be merged together", res["area"])

	// calculate the position of each image in the final canvas
	positions := calculatePositions(rect, images, *thumb_width, *thumb_height)

	// give a name to the canvas file and prepare it
	canvas_filename := filepath.Join(*dest_dir, randStr(20)+".jpg")
	canvas_image := image.NewRGBA(image.Rect(0, 0, *thumb_width*res["base"], *thumb_height*res["height"]))

	// iterate through the images, resize if necessary, decode and add to the canvas
	for _, image_path := range images {
		thumb := NewThumb(
			*thumb_width,
			*thumb_height,
			image_path,
		)
		thumb.SetDimension()
		img, err := thumb.DecodeIt()

		if err != nil {
			log.Printf("it was not possible to decode the image %s: %v", image_path, err)
		} else {
			x := positions[image_path][0]
			y := positions[image_path][1]
			draw.Draw(canvas_image, canvas_image.Bounds(), img, image.Point{x, y}, draw.Src)
		}
	}
	toimg, _ := os.Create(canvas_filename)
	defer toimg.Close()
	jpeg.Encode(toimg, canvas_image, &jpeg.Options{jpeg.DefaultQuality})

	log.Printf("canvas %s succesfully created", canvas_filename)
}
