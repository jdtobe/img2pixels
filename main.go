package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	fWidth := flag.Int("width", 16, "width of output pixels.")
	fHeight := flag.Int("height", 16, "height of output pixels.")
	flag.Parse()

	f, _ := os.Open(os.Args[1])
	originalImg, _, _ := image.Decode(f)
	f.Close()

	img := imaging.Resize(originalImg, *fWidth, *fHeight, imaging.Lanczos)
	// img := imaging.Resize(originalImg, *fWidth, *fHeight, imaging.CatmullRom)
	// img := imaging.Resize(originalImg, *fWidth, *fHeight, imaging.NearestNeighbor)

	for y := *fHeight - 1; y >= 0; y-- {
		if y%2 != 0 {
			for x := 0; x < *fWidth; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				fmt.Printf("0x%02X%02X%02X, ", r/256, g/256, b/256)
			}
		} else {
			for x := *fWidth - 1; x >= 0; x-- {
				r, g, b, _ := img.At(x, y).RGBA()
				fmt.Printf("0x%02X%02X%02X, ", r/256, g/256, b/256)
			}
		}
		fmt.Println()
	}

	f, err := os.Create("test.png")
	if err != nil {
		log.Fatalf("Error opening output file: %s", err)
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatalf("Error encoding output file: %s", err)
	}
}
