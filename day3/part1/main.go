package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func getLines(input_file string) ([]string, error) {
	file, err := os.Open(input_file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func Solve(input_file string) error {
	lines, err := getLines(input_file)
	if err != nil {
		return err
	}

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var problems []string
	for _, line := range lines {
		problems = slices.Concat(problems, r.FindAllString(line, -1))
	}

	total := 0
	for _, problem := range problems {
		parts := strings.Split(problem, ",")
		left, _ := strconv.Atoi(parts[0][4:])
		right, _ := strconv.Atoi(parts[1][:len(parts[1])-1])

		total += left * right
	}

	fmt.Println(total)

	return nil
}
