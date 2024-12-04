package main

import (
	"log"
	"os"

	"github.com/lafferjm/day2/part1"
	"github.com/lafferjm/day2/part2"
)

func main() {
	// Run part 1
	if err := part1.Solve("./input.txt"); err != nil {
		log.Printf("Error in part 1: %v", err)
		os.Exit(1)
	}

	// Run part 2
	if err := part2.Solve("./input.txt"); err != nil {
		log.Printf("Error in part 2: %v", err)
		os.Exit(1)
	}
}
