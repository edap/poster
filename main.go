package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	//"math"
	"os"
	"path/filepath"
	"regexp"
	//errors
)

// creare il test per countFiles
// far si che count files diventi count images chiama la conversione a jpg (o converta tutto a jpg)
// visto che ha i bytes, potrebbe dare un messaggio interessante durante il procedimento
// se non ci sono immagini, countImages mi dice di dargli una cartella con le immagini, e mi fa una faccina buffa

func main() {
	var (
		thumb_width  = flag.Int("thumb_width", 120, "the width of a single thumb")
		thumb_height = flag.Int("thumb_height", 90, "the height of a single thumb")
		source_dir   = flag.String("source_dir", "/home/da/to_merge", "the origin directory that contains the images to compose the grid")
		canvas_file  = rand_str(20) + ".jpg"
	)
	flag.Parse()
	fmt.Println(canvas_file)
	total_images := countImages(*source_dir)
	//preparing the empty grid
	res := map[string]int{
		"area":    total_images,
		"height":  0,
		"base":    0,
		"skipped": 0,
	}
	// fulfill te grid
	rect := CalculateRectangle(res)
	fmt.Println(rect)

	// CreateCanvas(
	// 	rect["height"],
	// 	rect["base"],
	// 	*thumb_width,
	// 	*thumb_height,
	// 	"/home/da/to_merge/merged/test.jpg")
	// decodificare sempre in jpg
	// convertToPNG converts from any recognized format to PNG.
	// func convertToPNG(w io.Writer, r io.Reader) error {
	//  img, _, err := image.Decode(r)
	//  if err != nil {
	//   return err
	//  }
	//  return png.Encode(w, img)
	// }

	// iterare tra le immagini e cambiare le coordinate
	// creare l'immagine di destinazione con un nome random
	// log, cosa succede se lancio il programma in una cartella con solo txt files?
	// prompt interattivo che mi dice: ci impiegherai circa x sec. Vuoi continuare. Oppure senza prompt

	// guardare come gli altri scrivono la documentazione, documentare una funzione ogni volta che la scrivi

	// create the canvas and coordinate
	back := image.NewRGBA(image.Rect(0, 0, 800, 600))
	var x, y = 0, 0

	files, _ := ioutil.ReadDir(*source_dir)
	for _, imgFile := range files {
		if isImage(imgFile.Name()) {

			img_name := imgFile.Name()
			src_path := filepath.Join(*source_dir, img_name)
			//try to open it
			if reader, err := os.Open(src_path); err == nil {
				defer reader.Close()

				im, _, err := image.DecodeConfig(reader)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
					continue
				}
				//check dimension and format
				thumb := NewThumb(im.Width, im.Height, *thumb_width, *thumb_height, img_name)
				thumb.HasDesiredDimension()
				//thumb.forceJpg

				//add to canvas
				img, _, _ := image.Decode(reader)
				x -= 180
				draw.Draw(back, back.Bounds(), img, image.Point{x, y}, draw.Src)

				fmt.Printf("%s %d %d\n", imgFile.Name(), thumb.CurrentWidth(), thumb.CurrentHeight())
			} else {
				log.Fatal("impossible to create target directory")
				fmt.Println("no")
			}
		}
	}

	toimg, _ := os.Create("/home/da/to_merge/" + canvas_file)
	defer toimg.Close()
	jpeg.Encode(toimg, back, &jpeg.Options{jpeg.DefaultQuality})

	// now go to the merged folder, count the images
	// use the composer
	// create an empty square
	// past the images in the empty square
}

func isImage(filename string) bool {
	is_img, _ := regexp.MatchString("png$|jpg$|jpeg$|gif$", filename)
	return is_img
}

// implementare log, o almeno, avere una politica coerente sugli errori
func countImages(source_dir string) int {
	dirname := source_dir
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tot := 0
	for _, fi := range fi {
		if fi.Mode().IsRegular() && isImage(fi.Name()) {
			tot += 1
			//fmt.Println(fi.Name(), fi.Size(), "bytes")
		}
	}
	return tot
}

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
