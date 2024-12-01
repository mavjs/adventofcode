package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type CoOrd struct {
	x, y int
}

func solution(filePath string) (int, int) {
	var part1, part2 = 0, 0

	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	fileScanner := bufio.NewScanner(data)

	partNum := make(map[CoOrd]int)
	symbols := make(map[CoOrd]string)
	//gears := []CoOrd{}

	row := 0
	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		// increase row counter
		row = row + 1
		for idx, char := range lineText {
			if char >= '0' && char <= '9' {
				partNum[CoOrd{idx, row}] = int(char) - '0'
			}

			if char != '.' {
				symbols[CoOrd{idx, row}] = string(char)
			}
		}
	}

	for idx, value := range partNum {
		fmt.Println(idx.x-1, idx.x+len(strconv.Itoa(value)))
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
