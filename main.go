package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc2021/day01"
	"aoc2021/day02"
	"aoc2021/day03"
)

var days = []func(){
	func() {},
	day01.Solve,
	day02.Solve,
	day03.Solve,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <day>", os.Args[0])
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Usage: %s <day>\n", os.Args[0])
		os.Exit(1)
	}

	if len(days)-1 < day || day < 1 {
		fmt.Printf("Valid day range: 1 - %d (inclusive)", len(days)-1)
		os.Exit(1)
	}

	days[day]()
}
