package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

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

	// be sure that at least 2 images are present
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
	rect := CalculateRectangle(res)
	log.Printf("%d images will be skipped", res["skipped"])
	log.Printf("%d images will be merged together", res["area"])
	// calculate the position of each image in the final canvas
	positions := CalculatePositions(rect, images, *thumb_width, *thumb_height)

	// give a name to the canvas file
	//filepath.Join(source_dir, fi.Name()
	canvas_file := filepath.Join(*dest_dir, "/"+randStr(20)+".jpg")
	// create the destination file
	back := image.NewRGBA(image.Rect(0, 0, *thumb_width*res["base"], *thumb_height*res["height"]))
	fmt.Println(canvas_file)

	// iterate through the images, resize if necessary, decode and add to the canvas
	for _, value := range images {
		// altezza e larghezza attuali dell'immagine, possono essere ricavati dopo, nella thumb
		thumb := NewThumb(
			120,
			90,
			*thumb_width,
			*thumb_height,
			value,
		)
		// SEI ARRIVATO QUI:
		// unit test su questa funzione, che deve ritornare il tipo di immagine
		thumb.GetFormatFromExtension()
		// se e' png forza a jpg. Se e' troppo complicato, lavoriamo solo con jpg all'inizio
		// mettere la cartella corrente come default, fare qualche prova
		img, err := thumb.DecodeIt()

		// img_file, _ := os.Open(value)
		// defer img_file.Close()
		// img, _, _ := image.Decode(img_file)

		if err != nil {
			log.Printf("it was not possible to decode the image:", err)
		}
		x := positions[value][0]
		y := positions[value][1]
		draw.Draw(back, back.Bounds(), img, image.Point{x, y}, draw.Src)
	}
	toimg, _ := os.Create(canvas_file)
	defer toimg.Close()
	jpeg.Encode(toimg, back, &jpeg.Options{jpeg.DefaultQuality})

	log.Printf("canvas %s succesfully created", canvas_file)
}

//DA FARE
// creare il test per listFiles
// far si che count files diventi count images chiama la conversione a jpg (o converta tutto a jpg)
// visto che ha i bytes, potrebbe dare un messaggio interessante durante il procedimento
// se non ci sono immagini, countImages mi dice di dargli una cartella con le immagini, e mi fa una faccina buffa
// log, cosa succede se lancio il programma in una cartella con solo txt files?
// prompt interattivo che mi dice: ci impiegherai circa x sec. Vuoi continuare. Oppure senza prompt
