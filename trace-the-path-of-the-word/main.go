package main

import (
	"fmt"
	"log"
)

func main() {

	result := traceWordPath("BISCUIT", [][]string{
		{"B", "I", "T", "R"},
		{"I", "U", "A", "S"},
		{"S", "C", "V", "W"},
		{"D", "O", "N", "E"},
	})
	if result == nil {
		log.Fatal("Not present")
	}

	fmt.Println(result)
}

func traceWordPath(word string, grid [][]string) [][]int {
	rows := len(grid)
	cols := len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if grid[r][c] == string(word[0]) {
				resPath := [][]int{}
				visited := make([][]bool, rows)
				for i := range visited {
					visited[i] = make([]bool, cols)
				}

				if search(grid, word, r, c, 0, visited, &resPath) {
					return resPath
				}
			}
		}
	}

	return nil
}

func search(grid [][]string, word string, r, c, currIndex int, visited [][]bool, resPath *[][]int) bool {
	if currIndex == len(word) {
		return true
	}

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}
	if visited[r][c] || grid[r][c] != string(word[currIndex]) {
		return false
	}

	visited[r][c] = true
	*resPath = append(*resPath, []int{r, c})

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, d := range dirs {
		if search(grid, word, r+d[0], c+d[1], currIndex+1, visited, resPath) {
			return true
		}
	}

	visited[r][c] = false
	*resPath = (*resPath)[:len(*resPath)-1]

	return false
}
