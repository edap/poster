package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	//"log"
	//"math"

	//errors
)

func main() {
	var (
		thumb_width  = flag.Int("thumb_width", 120, "the width of a single thumb")
		thumb_height = flag.Int("thumb_height", 90, "the height of a single thumb")
		source_dir   = flag.String("source_dir", "/home/da/to_merge", "the origin directory that contains the images to compose the grid")
		canvas_file  = *source_dir + "/" + randStr(20) + ".jpg"
	)
	flag.Parse()

	tot, images := listFiles(*source_dir)
	// procedi solo se a zero

	//preparing the empty grid
	res := map[string]int{"area": tot, "height": 0, "base": 0, "skipped": 0}
	rect := CalculateRectangle(res)
	positions := CalculatePositions(rect, images, *thumb_width, *thumb_height)

	fmt.Println(rect)
	fmt.Println(positions)

	// create the canvas and coordinate
	back := image.NewRGBA(image.Rect(0, 0, *thumb_width*res["base"], *thumb_height*res["height"]))

	for _, value := range images {
		img_file, _ := os.Open(value)
		defer img_file.Close()
		img, _, _ := image.Decode(img_file)

		x := positions[value][0]
		y := positions[value][1]
		draw.Draw(back, back.Bounds(), img, image.Point{x, y}, draw.Src)
	}
	toimg, _ := os.Create(canvas_file)
	defer toimg.Close()
	jpeg.Encode(toimg, back, &jpeg.Options{jpeg.DefaultQuality})
}

//DA FARE
// creare il test per listFiles
// far si che count files diventi count images chiama la conversione a jpg (o converta tutto a jpg)
// visto che ha i bytes, potrebbe dare un messaggio interessante durante il procedimento
// se non ci sono immagini, countImages mi dice di dargli una cartella con le immagini, e mi fa una faccina buffa
// log, cosa succede se lancio il programma in una cartella con solo txt files?
// prompt interattivo che mi dice: ci impiegherai circa x sec. Vuoi continuare. Oppure senza prompt
