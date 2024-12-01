package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// Numbers in Decimel format: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
var numberDecimel = []int{
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
}

func Part1(filePath string) int {
	var result = 0

	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	fileScanner := bufio.NewScanner(data)

	var numPairs = make([]int, 0)

	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		var numOnly = make([]int, 0)
		for _, char := range lineText {
			if slices.Contains(numberDecimel, int(char)) {
				numOnly = append(numOnly, (int(char) - '0'))
			}
		}

		result := 10*numOnly[0] + numOnly[len(numOnly)-1]
		numPairs = append(numPairs, result)

	}

	for _, pair := range numPairs {
		result = result + pair
	}
	return result
}

func Part2(filePath string) int {
	var result = 0
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	part_1 := Part1(filePath)
	part_2 := Part2(filePath)
	fmt.Println(part_1)
	fmt.Println(part_2)
}
