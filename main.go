package main

import (
	"fmt"
)

const defaultX = 90
const defaultY = 90
const defaultN = 50
const defaultOccurence = 0.05
const scale = 15

func printGrid(grid [][]bool) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	iterations := Generate(defaultOccurence, defaultX, defaultY, defaultN)
	SaveAsGif("test.gif", iterations, scale)
}
