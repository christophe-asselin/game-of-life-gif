package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
)

const delay = 10

func SaveAsGif(filename string, frames [][][]bool) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Could not create file " + filename)
	}
	defer f.Close()

	err = gif.EncodeAll(f, framesToGif(frames))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func framesToGif(frames [][][]bool) *gif.GIF {
	images := make([]*image.Paletted, len(frames))
	delays := make([]int, len(frames))
	disposal := make([]byte, len(frames))
	for i, frame := range frames {
		images[i] = gridToPaletted(frame)
		delays[i] = delay
		disposal[i] = gif.DisposalNone
	}

	width := len(frames[0][0])
	height := len(frames[0])

	config := image.Config{
		ColorModel: nil,
		Width: width,
		Height: height,
	}

	return &gif.GIF{
		Image: images,
		Delay: delays,
		LoopCount: 0,
		Disposal: disposal,
		Config: config,
		BackgroundIndex: 0,
	}
}

func gridToPaletted(grid [][]bool) *image.Paletted {
	if len(grid) == 0 {
		log.Fatal("Grid is empty")
	}

	width := len(grid[0])
	height := len(grid)

	palette := []color.Color{color.Black, color.White}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)	
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	return img
}

