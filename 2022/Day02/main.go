package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(filePath string) (int, int) {
	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	fileScanner := bufio.NewScanner(data)

	total := 0

	for fileScanner.Scan() {
		score := 0
		lineText := fileScanner.Text()

		splitText := strings.Split(lineText, " ")
		opponent := splitText[0]
		move := splitText[1]

		// Part 1: X = Rock, Y = Paper, Z = Scissor, Part 2: X = lose, Y = draw, Z = win
		if move == "X" {
			score = score + 1
		} else if move == "Y" {
			score = score + 2
		} else {
			score = score + 3
		}

		// A = Rock, B = Paper, C = Scissor
		if (opponent == "A" && move == "X") || (opponent == "B" && move == "Y") || (opponent == "C" && move == "Z") {
			score = score + 3
		}

		if (opponent == "A" && move == "Y") || (opponent == "B" && move == "Z") || (opponent == "C" && move == "X") {
			score = score + 6
		}

		if (opponent == "A" && move == "Z") || (opponent == "B" && move == "X") || (opponent == "C" && move == "Y") {
			score = score + 0
		}

		total = total + score
	}

	return total, total
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	part_1, part_2 := solution(filePath)
	fmt.Println("total score:", part_1)
	fmt.Println("total score2:", part_2)

}
