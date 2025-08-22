package main

import (
	"flag"
	"log"
)

var occurence = flag.Float64("p", 0.05, "Probability of a cell being alive in the initial state, must be in range [0, 1]")
var x = flag.Int("x", 90, "Width of image in number of cells")
var y = flag.Int("y", 90, "Height of image in number of cells")
var n = flag.Int("n", 30, "Number of iterations")
var scale = flag.Int("s", 15, "Scale factor for cells, e.g. scale of 1 means each cell will be 1 pixel wide")
var filename = flag.String("o", "gameoflife.gif", "Output file name")

func main() {
	flag.Parse()

	if *occurence < 0 || *occurence > 1 {
		log.Fatal("p must be in range [0, 1]")
	}
	if *x <= 0 {
		log.Fatal("x must be at least 1")
	}
	if *y <= 0 {
		log.Fatal("y must be at least 1")
	}
	if *n <= 0 || *n > 100 {
		log.Fatal("n must be between 1 and 100")
	}
	if *scale <= 0 || *scale > 100 {
		log.Fatal("s must be between 1 and 100")
	}
	if len(*filename) == 0 {
		log.Fatal("o must be non empty")
	}

	frames := Generate(*occurence, *x, *y, *n)
	SaveAsGif(*filename, frames, *scale)
}
