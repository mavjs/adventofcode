package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func md5sum(key string, i int) string {
	hash := md5.New()
	hash.Write([]byte(key + strconv.Itoa(i)))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	input := "abcdef"
	zeroes := 5

	if len(os.Args) > 2 {
		input = os.Args[1]
		zeroes, _ = strconv.Atoi(os.Args[2])
	}

	fmt.Println("input:", input, "zeroes:", zeroes)
	target1 := strings.Repeat("0", zeroes)
	counter, _ := strconv.Atoi(fmt.Sprintf("1%s0", target1))

	for idx := 0; idx < counter; idx++ {
		result := md5sum(input, idx)
		if result[:zeroes] == target1 {
			fmt.Println("[Part 1]: ", idx)
			break
		}
	}

	target2 := strings.Repeat("0", zeroes+1)
	counter2, _ := strconv.Atoi(fmt.Sprintf("1%s0", target2))
	for idx := 0; idx < counter2; idx++ {
		result := md5sum(input, idx)
		if result[:zeroes+1] == target2 {
			fmt.Println("[Part 2]: ", idx)
			break
		}
	}
}
