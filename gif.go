package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"os"
)

const delay = 10

func SaveAsGif(filename string, frames [][][]bool, scale int) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Could not create file " + filename)
	}
	defer f.Close()

	err = gif.EncodeAll(f, framesToGif(frames, scale))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func framesToGif(frames [][][]bool, scale int) *gif.GIF {
	images := make([]*image.Paletted, len(frames))
	delays := make([]int, len(frames))
	disposal := make([]byte, len(frames))
	for i, frame := range frames {
		images[i] = gridToPaletted(frame, scale)
		delays[i] = delay
		disposal[i] = gif.DisposalNone
	}

	width := len(frames[0][0]) * scale
	height := len(frames[0]) * scale

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

func gridToPaletted(grid [][]bool, scale int) *image.Paletted {
	if len(grid) == 0 {
		log.Fatal("Grid is empty")
	}

	width := len(grid[0]) * scale
	height := len(grid) * scale

	palette := []color.Color{color.Black, color.White}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)	
	for y := range grid {
		for x := range grid[y] {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[y][x] {
						img.Set(x * scale + i, y * scale + j, color.Black)
					} else {
						img.Set(x * scale + i, y * scale + j, color.White)
					}
				}
			}
		}
	}

	return img
}

