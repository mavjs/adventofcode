package main

import (
	"bufio"
	"fmt"
	"os"
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

func scoreCalulation(inputChar string) int {
	retVal := 0

	// Turn string into run and then convert to int to get ASCII number
	ascii_decimal := int([]rune(inputChar)[0])

	// a-z = 1-16, A-Z = 27-52
	// a = 97, thus, 97 - 96 = 1
	// A = 65, thus 65 - 38 = 27
	if strings.ToUpper(inputChar) == inputChar {
		score := ascii_decimal - 38
		retVal = retVal + score
	} else {
		score := ascii_decimal - 96
		retVal = retVal + score
	}

	return retVal
}

func solution1(filePath string) int {

	fileLines := fileRead(filePath)

	total := 0
	for _, lineText := range fileLines {
		contents := strings.Split(lineText, "")

		mid_point := (len(contents)) / 2

		first_rucksack := contents[:mid_point]
		second_rucksack := contents[mid_point:]

		out := ""
		bucket := map[string]bool{}

		for _, val1 := range second_rucksack {
			for _, val2 := range first_rucksack {
				if val1 == val2 && !bucket[val1] {
					out = val1
					bucket[val1] = true
				}
			}
		}

		total = total + scoreCalulation(out)

	}

	return total
}

func solution2(filePath string) int {
	lines := fileRead(filePath)

	total := 0

	for i := 0; i < len(lines); i += 3 {
		elf1 := strings.Split(lines[i], "")
		elf2 := strings.Split(lines[i+1], "")
		elf3 := strings.Split(lines[i+2], "")

		out := ""
		bucket := map[string]bool{}

		for _, val1 := range elf1 {
			for _, val2 := range elf2 {
				for _, val3 := range elf3 {
					if val1 == val2 && val1 == val3 && !bucket[val3] {
						out = val3
						bucket[val3] = true
					}
				}
			}
		}

		total = total + scoreCalulation(out)
	}

	return total
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide an input file to process")
		os.Exit(-1)
	}
	filePath := os.Args[1]
	fmt.Println("part 1: ", solution1(filePath))
	fmt.Println("part 2: ", solution2(filePath))
}
