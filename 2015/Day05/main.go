package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func solution1(filePath string) (int, int) {
	part_1 := 0
	part_2 := 0

	vowels := "aeiou"
	ban_strings := regexp.MustCompile("(ab|cd|pq|xy)")

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	input := string(data)

	for _, line := range strings.Split(input, "\n") {
		vowel_count := 0
		for _, char := range line {
			if strings.ContainsRune(vowels, char) {
				vowel_count = vowel_count + 1
			}
		}

		doubleWord := false
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				doubleWord = true
				break
			}
		}

		if vowel_count >= 3 && !ban_strings.MatchString(line) && doubleWord {
			part_1 = part_1 + 1
		}
	}

	return part_1, part_2
}

func solution2(filePath string) (int, int) {
	part_1 := 0
	part_2 := 0

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	input := string(data)

	rule1 := func(line string) bool {
		for i := 0; i < len(line)-2; i++ {
			toMatch := line[i : i+2]
			for j := i + 2; j < len(line)-1; j++ {
				if line[j:j+2] == toMatch {
					return true
				}
			}
		}
		return false
	}

	rule2 := func(line string) bool {
		for x := 0; x < len(line)-2; x++ {
			if line[x] == line[x+2] {
				return true
			}
		}
		return false
	}
	for _, line := range strings.Split(input, "\n") {
		two_letter_pair := rule1(line)
		one_letter_repeat := rule2(line)

		if two_letter_pair && one_letter_repeat {
			part_1 = part_1 + 1
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
	part_1, _ := solution1(filePath)
	part_2, _ := solution2(filePath)
	fmt.Println("[part 1] nice strings:", part_1)
	fmt.Println("[part 2] nice strings:", part_2)

}
