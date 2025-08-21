package main

import (
	"fmt"
)

const defaultX = 1600
const defaultY = 900
const defaultN = 50
const defaultOccurence = 0.05

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
	SaveAsGif("test.gif", iterations)
}
