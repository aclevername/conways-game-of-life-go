package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
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

	for y, line := range lines {
		for x, c := range line {
			value := 0
			if c == '1' {
				value = 1
			}
			currentGrid[y][x] = value
		}
	}
	printGrid(currentGrid)
	for {
		updateGrid(currentGrid)
		printGrid(currentGrid)
		time.Sleep(time.Second / 10)
	}
}

func printGrid(grid [][]int) {
	fmt.Println()
	for _, row := range grid {
		for _, value := range row {
			if value == 0 {
				fmt.Print("-")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

func updateGrid(grid [][]int) {
	futureGrid := makeGrid(len(grid))
	for y, row := range grid {
		for x, _ := range row {
			futureGrid[y][x] = getValue(grid, x, y)
		}
	}

	for y, row := range grid {
		for x, _ := range row {
			grid[y][x] = futureGrid[y][x]
		}
	}
}

type position struct {
	x int
	y int
}

func getValue(grid [][]int, x, y int) int {
	cells := []position{
		position{x + 1, y - 1}, position{x + 1, y}, position{x + 1, y + 1},
		position{x, y - 1}, position{x, y + 1},
		position{x - 1, y - 1}, position{x - 1, y}, position{x - 1, y + 1},
	}

	count := 0
	size := len(grid)

	for _, cell := range cells {
		x := cell.x
		y := cell.y
		if x >= 0 && x < size && y >= 0 && y < size && grid[cell.y][cell.x] == 1 {
			count++
		}
	}

	if grid[y][x] == 1 {
		if count == 2 || count == 3 {
			return 1
		}
	} else {
		if count == 3 {
			return 1
		}
	}
	return 0
}

func makeGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}
