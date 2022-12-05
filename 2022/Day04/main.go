package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fileRead(filePath string) []string {
	lines := make([]string, 0)

	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(data)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

func solution(filePath string) (int, int) {

	fileLines := fileRead(filePath)

	// totals for each parts of solution
	part1_total := 0
	part2_total := 0

	for _, lineText := range fileLines {
		contents := strings.Split(lineText, ",")

		cleaning_sector := make([][]int, 0)

		for _, elf := range contents {
			id_num := make([]int, 0)
			sections := strings.Split(elf, "-")

			section1, _ := strconv.Atoi(sections[0])
			section2, _ := strconv.Atoi(sections[1])
			for i := section1; i <= section2; i++ {
				id_num = append(id_num, i)
			}

			cleaning_sector = append(cleaning_sector, id_num)
		}

		out := make([]int, 0)
		bucket := map[int]bool{}

		sector_1 := cleaning_sector[0]
		sector_2 := cleaning_sector[1]
		for _, val1 := range sector_1 {
			for _, val2 := range sector_2 {
				if val1 == val2 && !bucket[val1] {
					out = append(out, val1)
					bucket[val1] = true
				}
			}
		}

		if out != nil && (len(out) == len(sector_1) || len(out) == len(sector_2)) {
			part1_total = part1_total + 1
		}

		if out != nil && len(out) > 0 {
			part2_total = part2_total + 1
		}
	}

	return part1_total, part2_total
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	part_1, part_2 := solution(filePath)
	fmt.Println("part 1: ", part_1)
	fmt.Println("part 2: ", part_2)
}
