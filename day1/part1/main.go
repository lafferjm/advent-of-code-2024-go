package part1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve(input_file string) {
	left := []int{}
	right := []int{}

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

		rightValue, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Failed to convert right value")
			os.Exit(1)
		}
		right = append(right, rightValue)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := math.Abs(float64(left[i] - right[i]))
		totalDistance = totalDistance + int(distance)
	}

	fmt.Println(totalDistance)
}
