package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(input_file string) {
	var left []int
	var right = make(map[string]int)

	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		leftValue, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Failed to convert left value")
			os.Exit(1)
		}
		left = append(left, leftValue)

		right[values[1]] += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	similarityScore := 0
	for _, val := range left {
		similarityScore = similarityScore + (val * right[strconv.Itoa(val)])
	}

	fmt.Println(similarityScore)
}
