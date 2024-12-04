package part1

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Println(lines)

	return nil
}
