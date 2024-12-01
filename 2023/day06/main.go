package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	time := make([]string, 0)
	distance := make([]string, 0)

	inputStr := make([]string, 0)

	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		_, input, _ := strings.Cut(lineText, ": ")
		inputStr = append(inputStr, input)
	}

	time = append(time, strings.Fields(inputStr[0])...)

	distance = append(distance, strings.Fields(inputStr[1])...)

	records := make([]int, 0)
	for idx, ms := range time {
		counter := 0
		maxDistance, _ := strconv.Atoi(distance[idx])
		maxMS, _ := strconv.Atoi(ms)
		for i := 1; i <= maxMS; i++ {
			if i*(maxMS-i) > maxDistance {
				counter = counter + 1
			}
		}
		records = append(records, counter)
	}

	part1 = 1
	for _, record := range records {
		part1 = part1 * record
	}

	newTime, _ := strconv.Atoi(strings.Join(time, ""))
	newDistance, _ := strconv.Atoi(strings.Join(distance, ""))

	fmt.Println("New Time:", newTime)
	fmt.Println("New Distance:", newDistance)

	for i := 1; i < newTime; i++ {
		if i*(newTime-i) > newDistance {
			part2 = part2 + 1
		}
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
