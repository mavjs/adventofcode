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

	fileScanner := bufio.NewScanner(data)

	cardNum := 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		cardNum = cardNum + 1
		_, inputStr, _ := strings.Cut(lineText, ": ")
		winningNum, drawNum, _ := strings.Cut(inputStr, " | ")

		// Winning Number HashMap per line
		win := map[string]bool{}
		// Points counter per line
		points := 0

		for _, num := range strings.Fields(winningNum) {
			win[num] = true
		}

		for _, num := range strings.Fields(drawNum) {
			if win[num] {
				points = points + 1
			}
		}
		points = (1 << points) >> 1

		part1 = part1 + points
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
