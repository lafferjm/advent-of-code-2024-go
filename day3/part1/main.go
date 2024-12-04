package part1

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

func buildBuffer(contents string) string {
	var buffer bytes.Buffer

	validCharacters := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ")", ","}

	for i := 0; i < len(contents); i++ {
		if string(contents[i]) == ")" {
			break
		} else if slices.Contains(validCharacters, string(contents[i])) {
			buffer.WriteString(string(contents[i]))
		} else {
			return ""
		}
	}

	return buffer.String()
}

func Solve(input_file string) error {
	lines, err := getLines(input_file)
	if err != nil {
		return err
	}

	var problems []string
	for _, line := range lines {
		for i := 0; i < len(line)-4; i++ {
			if line[i:i+4] == "mul(" {
				buffer := buildBuffer(line[i+4:])
				problems = append(problems, buffer)
			}
		}
	}

	total := 0
	for _, problem := range problems {
		if len(problem) > 0 {
			parts := strings.Split(problem, ",")
			first, err := strconv.Atoi(parts[0])
			if err != nil {
				return err
			}

			second, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}

			total += first * second
		}
	}

	fmt.Println(total)

	return nil
}
