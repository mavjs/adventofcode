package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(filePath string) (int, int) {
	part_1 := 0
	part_2 := 0

	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(data)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		text := strings.Split(line, "x")

		inputs := make([]int, len(text))
		for idx, value := range text {
			inputs[idx], _ = strconv.Atoi(value)
		}

		area := (2 * inputs[0] * inputs[1]) + (2 * inputs[1] * inputs[2]) + (2 * inputs[2] * inputs[0])
		sort.Ints(inputs)
		area = area + (inputs[0] * inputs[1])
		part_1 = part_1 + area

		length := (2 * inputs[0]) + (2 * inputs[1]) + (inputs[0] * inputs[1] * inputs[2])
		part_2 = part_2 + length

	}

	return part_1, part_2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	part_1, part_2 := solution(filePath)
	fmt.Println("[part 1]:", part_1)
	fmt.Println("[part 2]:", part_2)
}
