package main

import (
	"fmt"
	"os"
	"strings"
)

func solution(input string) (int, int) {
	part_1 := 0
	part_2 := 0

	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	grid1 := make([][]int, 1000)
	for i := range grid1 {
		grid1[i] = make([]int, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "turn on"):
			var x, y, x1, y1 int
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &x, &y, &x1, &y1)
			for i := x; i <= x1; i++ {
				for j := y; j <= y1; j++ {
					grid[i][j] = true
					grid1[i][j]++
				}
			}
		case strings.HasPrefix(line, "toggle"):
			var x, y, x1, y1 int
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x, &y, &x1, &y1)
			for i := x; i <= x1; i++ {
				for j := y; j <= y1; j++ {
					grid[i][j] = !grid[i][j]
					grid1[i][j] = grid1[i][j] + 2
				}
			}
		case strings.HasPrefix(line, "turn off"):
			var x, y, x1, y1 int
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &x, &y, &x1, &y1)
			for i := x; i <= x1; i++ {
				for j := y; j <= y1; j++ {
					grid[i][j] = false
					if grid1[i][j] > 0 {
						grid1[i][j]--
					}
				}
			}
		default:
			panic("unhandled!")
		}
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] {
				part_1 = part_1 + 1
			}
		}
	}

	for x := range grid1 {
		for y := range grid1[x] {
			part_2 = part_2 + grid1[x][y]
		}
	}

	return part_1, part_2
}

func main() {
	part_1 := 0
	part_2 := 0

	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	input := string(data)

	part_1, part_2 = solution(input)

	fmt.Println("[Part 1]:", part_1)
	fmt.Println("[Part 2]:", part_2)
}
