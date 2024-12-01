package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculatePrevNumber(s []int) (r int) {
	var finish bool = true
	ss := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		ss[i] = s[i+1] - s[i]
		if ss[i] != 0 {
			finish = false
		}
	}
	if finish {
		return s[0]
	}
	return s[0] - calculatePrevNumber(ss)
}

func calculateNextNumber(s []int) (r int) {
	var finish bool = true
	ss := make([]int, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		ss[i] = s[i+1] - s[i]
		if ss[i] != 0 {
			finish = false
		}
	}
	if finish {
		return s[0]
	}
	return s[len(s)-1] + calculateNextNumber(ss)
}

func solution(filePath string) (int, int) {
	var part1, part2 = 0, 0

	data, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	fileScanner := bufio.NewScanner(data)

	sensorReports := make([][]int, 0)
	rows := 0
	for fileScanner.Scan() {
		rows = rows + 1
		numSlice := make([]int, 0)
		for _, num := range strings.Fields(fileScanner.Text()) {
			numInt, _ := strconv.Atoi(num)
			numSlice = append(numSlice, numInt)
		}

		sensorReports = append(sensorReports, numSlice)
	}

	for _, row := range sensorReports {
		part1 = part1 + calculateNextNumber(row)
		part2 = part2 + calculatePrevNumber(row)
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
