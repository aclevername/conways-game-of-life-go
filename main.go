package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("infinite-pattern")
	if err != nil {
		fmt.Print(err)
	}

	lines := strings.Split(string(fileContent), "\n")
	fmt.Print(lines)

	size := len(lines[0])

	currentGrid := makeGrid(size)
	futureGrid := makeGrid(size)

	for y, line := range lines {
		for x, c := range line {
			value := 0
			if c == '1' {
				value = 1
			}
			fmt.Print(value)
			currentGrid[y][x] = value
		}
}

func makeGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}
