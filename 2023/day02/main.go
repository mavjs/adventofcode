package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(filePath string) (int, int) {
	var part1, part2 = 0, 0

	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	var limit = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	fileScanner := bufio.NewScanner(data)

	var gameId = 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		// increate gameId counter
		gameId = gameId + 1

		_, gameData, _ := strings.Cut(lineText, ":")
		gameSets := strings.Split(gameData, ";")

		invalid := false
		properGameSets := make(map[string]int)
		for _, gameSet := range gameSets {
			cubes := strings.Split(gameSet, ",")
			for _, cube := range cubes {
				var cubeNum int
				var cubeColor string

				fmt.Sscanf(cube, "%d %s", &cubeNum, &cubeColor)

				properGameSets[cubeColor] = max(properGameSets[cubeColor], cubeNum)

				if cubeNum > limit[cubeColor] {
					invalid = true
				}
			}
		}

		if !invalid {
			part1 = part1 + gameId
		}

		part2 = part2 + (properGameSets["red"] * properGameSets["green"] * properGameSets["blue"])
	}

	return part1, part2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	part_1, part_2 := solution(filePath)
	fmt.Println(part_1)
	fmt.Println(part_2)
}
