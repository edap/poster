package main

import (
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	//"math"
	"os"
	"path/filepath"
	"regexp"
	//errors
)

// creare il test per countFiles
// far si che count files diventi count images, e che chekki il tipo di immagine(o converta tutto a jpg)
// visto che ha i bytes, potrebbe dare un messaggio interessante durante il procedimento
// se non ci sono immagini, countImages mi dice di dargli una cartella con le immagini, e mi fa una faccina buffa

func main() {
	var (
		//ar allow_scaling = scale images if it is a prime number
		thumb_width  = flag.Int("thumb_width", 120, "the width of a single thumb")
		thumb_height = flag.Int("thumb_height", 90, "the height of a single thumb")
		//erase_original = flag.Bool("erase_original", false, "erase the original thumbs after being merged into the grid")
		source_dir = flag.String("source_dir", "/home/da/to_merge", "the origin directory that contains the images to compose the grid")
		target_dir = flag.String("target_dir", "/home/da/to_merge/merged", "the destination directory that will contain the final grid")
	)
	flag.Parse()

	if err := os.MkdirAll(*target_dir, 0755); err != nil {
		log.Fatal("impossible to create target directory")
	}
	total_images := countImages(*source_dir)

	res := map[string]int{
		"area":    total_images,
		"height":  0,
		"base":    0,
		"skipped": 0,
	}
	rect := CalculateRectangle(res)
	fmt.Println(rect)
	CreateCanvas(
		rect["height"],
		rect["base"],
		*thumb_width,
		*thumb_height,
		"/home/da/to_merge/merged/test.jpg")
	// fin qua ci siamo, la destinazione e' creata.
	// ORA
	// copiare una thumb in una posizione desiderata
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
	// rimuovere l'opzione delete originals
	// rimuovere l'opzione dest source, semmai metterci dest file. default la current directory
	// log, cosa succede se lancio il programma in una cartella con solo txt files?
	// prompt interattivo che mi dice: ci impiegherai circa x sec. Vuoi continuare. Oppure senza prompt
	// scrivere tests
	// guardare come gli altri scrivono la documentazione, documentare una funzione ogni volta che la scrivi
	// impararti i metodi principali di os e di image

	files, _ := ioutil.ReadDir(*source_dir)
	for _, imgFile := range files {
		if isImage(imgFile.Name()) {

			img_name := imgFile.Name()
			src_path := filepath.Join(*source_dir, img_name)
			//dst_path := filepath.Join(*target_dir, img_name)
			if reader, err := os.Open(src_path); err == nil {
				defer reader.Close()

				im, _, err := image.DecodeConfig(reader)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
					continue
				}
				fmt.Println(im.Width)
				fmt.Println(im.Height)
				fmt.Println(*thumb_width)
				fmt.Println(*thumb_height)
				fmt.Println(img_name)
				// qui va implementata la logica che conta le immagini, calcola un quadro,
				// e dice su quante linee devono essere disposte le immagini

				// thumb := &Thumb{
				// 	width:          im.Width,
				// 	height:         im.Height,
				// 	desired_width:  *thumb_width,
				// 	desired_height: *thumb_height,
				// 	img_name:       img_name,
				// }

				thumb := NewThumb(im.Width, im.Height, *thumb_width, *thumb_height, img_name)
				thumb.HasDesiredDimension()

				// if *erase_original == true {
				// 	thumb.Move(src_path, dst_path)
				// } else {
				// 	thumb.Copy(src_path, dst_path)
				// }
				// thumb.Scale(dst_path)

				fmt.Printf("%s %d %d\n", imgFile.Name(), thumb.CurrentWidth(), thumb.CurrentHeight())
			} else {
				fmt.Println("no")
			}
		}
	}

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

// conta le immagini che ha nella directory, cerca di rispettare le proporzioni
// (date come parametro), e mette quello che ci sta. Quello che ci sta lo cancella (opzione delete?) quello che non ci sta viene salvato su un file, che fa da registro

// package main

// import (
// 	"fmt"
// 	"image"
// 	"image/draw"
// 	"image/jpeg"
// 	"os"
// )

// func main() {
// 	fImg1, _ := os.Open("/home/da/to_merge/Gwx_TPMT3Bg.jpg")
// 	defer fImg1.Close()
// 	img1, _, _ := image.Decode(fImg1)

// 	m := image.NewRGBA(image.Rect(0, 0, 800, 600))
// 	draw.Draw(m, m.Bounds(), img1, image.Point{0, 0}, draw.Src)
// 	//draw.Draw(m, m.Bounds(), img2, image.Point{-200,-200}, draw.Src)

// 	toimg, _ := os.Create("/home/da/to_merge/new.jpg")
// 	defer toimg.Close()

// 	jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})
// }

// http://golang.org/src/pkg/image/jpeg/reader.go?s=10946:10998#L343
// http://golang.org/src/pkg/image/jpeg/reader.go?s=10744:10789#L336
// func (t Thumb) Scale() (int, error) {
//   // if HasDesiredDimension
//   m := resize.Resize(300, 200, t.decoded_img, resize.NearestNeighbor)
//   out, _ := os.Create("test_resized.jpg")
//   // if err != nil {
//   //  log.Fatal(err)
//   // }
//   defer out.Close()

//   // write new image to file
//   jpeg.Encode(out, m, nil)
//   return 2, nil
// }
