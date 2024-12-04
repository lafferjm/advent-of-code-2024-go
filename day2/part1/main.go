package part1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Check if sequence is increasing or decreasing
	increasing := levels[1] > levels[0]
	
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		
		// Check if difference is between 1 and 3
		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}
		
		// Check if sequence maintains direction
		if increasing && diff < 0 {
			return false
		}
		if !increasing && diff > 0 {
			return false
		}
	}
	
	return true
}

// parseLevels converts a space-separated string of numbers into a slice of integers
func parseLevels(line string) ([]int, error) {
	if line == "" {
		return nil, nil
	}

	numStrs := strings.Fields(line)
	levels := make([]int, len(numStrs))
	
	for i, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("parsing number at position %d: %w", i, err)
		}
		levels[i] = num
	}

	return levels, nil
}

func Solve(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		levels, err := parseLevels(scanner.Text())
		if err != nil {
			return fmt.Errorf("parsing levels: %w", err)
		}
		if levels == nil {
			continue
		}

		if isReportSafe(levels) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	fmt.Printf("Part 1 - Number of safe reports: %d\n", safeCount)
	return nil
}
