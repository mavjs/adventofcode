package main

import (
	"fmt"
	"os"
)

func solution(filePath string) (int, int) {
	part_1 := 0
	part_2 := 0

	found := false

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	for idx, text := range data {
		if text == '(' {
			part_1 = part_1 + 1
		}
		if text == ')' {
			part_1 = part_1 - 1
		}

		if part_1 == -1 && !found {
			part_2 = idx + 1
			found = true
		}
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
	fmt.Println("[part 1] floors:", part_1)
	fmt.Println("[part 2] floors:", part_2)

}
