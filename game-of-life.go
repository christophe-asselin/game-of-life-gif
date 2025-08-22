package main

import (
	"math/rand"
)

func Generate(occurence float64, x, y, n int) [][][]bool {
	grid := createEmptyGrid(x, y)
	addRandomCells(grid, occurence)
	iterations := make([][][]bool, n - 1)
	iterations[0] = grid

	for i := 1; i < n - 1; i++ {
		grid = performIteration(grid)
		iterations[i] = grid
	}

	return iterations
}

func createEmptyGrid(x, y int) [][]bool {
	arr := make([][]bool, y)
	for i := range arr {
		arr[i] = make([]bool, x)
	}
	return arr
}

func addRandomCells(grid [][]bool, occurence float64) {
	for i := range grid {
		for j := range grid[i] {
			if randN := rand.Float64(); randN < occurence {
				grid[i][j] = true
			}
		}
	}
}

func countLiveNeighbors(grid [][]bool, i, j int) int {
	liveCount := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}

			u := i + di
			v := j + dj

			if u >= 0 && u < len(grid) && v >= 0 && v < len(grid[0]) {
				if grid[u][v] {
					liveCount++
				}
			}

		}
	}
	return liveCount
}

func performIteration(grid [][]bool) [][]bool {
	newArr := createEmptyGrid(len(grid[0]), len(grid))
	for i := range grid {
		for j, isAlive := range grid[i] {
			liveNeighbors := countLiveNeighbors(grid, i, j)
			willBeAlive := isAlive
			if isAlive {
				if liveNeighbors < 2 {
					willBeAlive = false // underpopulation
				} else if liveNeighbors > 3 {
					willBeAlive = false // overpopulation
				} else {
					willBeAlive = true // survives
				}
			} else {
				if liveNeighbors == 3 {
					willBeAlive = true // reproduction
				}
			}
			newArr[i][j] = willBeAlive	
		}
	}
	return newArr
}