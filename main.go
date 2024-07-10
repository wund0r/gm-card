package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	ROW_1     = 427
	COL_1     = 422
	ROW_SHIFT = 49
	COL_SHIFT = 81
)

// based-1
func getPoint(row, col int) image.Point {
	return image.Point{ROW_1 + (ROW_SHIFT * (row - 1)), COL_1 + (COL_SHIFT * (col - 1))}
}

func main() {
	log.Println("Hello world")
	reader, err := os.Open("./background.png")
	if err != nil {
		log.Fatal(err)
	}
	background, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	reader, err = os.Open("./mark.png")
	if err != nil {
		log.Fatal(err)
	}
	mark, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	output := image.NewRGBA(background.Bounds())

	draw.Draw(output, output.Bounds(), background, image.Point{0, 0}, draw.Src)
	position := getPoint(5, 6)
	draw.Draw(output, mark.Bounds().Add(position), mark, image.Point{0, 0}, draw.Over)

	outputFile, err := os.Create("./result.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outputFile, output)
	if err != nil {
		log.Fatal(err)
	}

}
