package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func markHouse(houses map[int]map[int]int, x, y int) {
	if houses == nil {
		houses = make(map[int]map[int]int)
	}

	if houses[x] == nil {
		houses[x] = make(map[int]int)
	}

	houses[x][y]++
}

func solution1(data string) (int, int) {
	part_1 := 0
	part_2 := 0

	houses := make(map[int]map[int]int)
	x, y := 0, 0

	markHouse(houses, x, y)

	for _, direction := range data {
		switch direction {
		case '^':
			y++
		case '>':
			x++
		case 'v':
			y--
		case '<':
			x--
		}
		markHouse(houses, x, y)
	}

	for _, house := range houses {
		part_1 = part_1 + len(house)
	}

	return part_1, part_2
}

func solution2(data string) (int, int) {
	part_1 := 0
	part_2 := 0

	houses := make(map[int]map[int]int)
	x, y := 0, 0
	x1, y1 := 0, 0

	markHouse(houses, x, y)
	markHouse(houses, x1, y1)

	for idx, direction := range data {
		if idx%2 == 0 {
			switch direction {
			case '^':
				y++
			case '>':
				x++
			case 'v':
				y--
			case '<':
				x--
			}
			markHouse(houses, x, y)
		} else {
			switch direction {
			case '^':
				y1++
			case '>':
				x1++
			case 'v':
				y1--
			case '<':
				x1--
			}
			markHouse(houses, x1, y1)
		}
	}

	for _, house := range houses {
		part_1 = part_1 + len(house)
	}

	return part_1, part_2
}

func main() {
	data, err := os.Open("data/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	testA_1, testA_2 := solution1(">")
	if testA_1 != 2 {
		log.Fatalln("[testA]", testA_1, testA_2)
	}

	testB_1, testB_2 := solution1("^>v<")
	if testB_1 != 4 {
		log.Fatalln("[testB]", testB_1, testB_2)
	}

	testC_1, testC_2 := solution1("^v^v^v^v^v")
	if testC_1 != 2 {
		log.Fatalln("[testC]", testC_1, testC_2)
	}

	testA1_1, testA1_2 := solution2("^v")
	if testA1_1 != 3 {
		log.Fatalln("[testA1]", testA1_1, testA1_2)
	}

	testB1_1, testB1_2 := solution2("^>v<")
	if testB1_1 != 3 {
		log.Fatalln("[testB1]", testB1_1, testB1_2)
	}

	testC1_1, testC1_2 := solution2("^v^v^v^v^v")
	if testC1_1 != 11 {
		log.Fatalln("[testC1]", testC1_1, testC1_2)
	}

	fileScanner := bufio.NewScanner(data)

	part_1 := 0
	part_2 := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		solution_1, _ := solution1(line)
		solution_2, _ := solution2(line)
		part_1 = part_1 + solution_1
		part_2 = part_2 + solution_2
	}

	fmt.Println("[part 1]:", part_1)
	fmt.Println("[part 2]:", part_2)
}
