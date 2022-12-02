package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solution(filePath string) (int, int) {
	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	fileScanner := bufio.NewScanner(data)

	calories := 0
	elves := make([]int, 0)

	for fileScanner.Scan() {
		lineText := fileScanner.Text()

		if lineText == "" {
			elves = append(elves, calories)
			calories = 0
			continue
		}

		calorie, err := strconv.Atoi(lineText)
		if err != nil {
			fmt.Println(err)
		}
		calories = calories + calorie
	}
	if calories > 0 {
		elves = append(elves, calories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	part_1 := elves[0]
	part_2 := elves[0] + elves[1] + elves[2]

	return part_1, part_2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	part_1, part_2 := solution(filePath)
	fmt.Println("most calories (1 elf):", part_1)
	fmt.Println("most calories (top 3 elves):", part_2)

}
