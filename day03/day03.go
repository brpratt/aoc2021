package day03

import (
	"aoc2021/util"
	"fmt"
)

func Solve() {
	lines, err := util.ReadLines("day03/input.txt")
	if err != nil {
		panic(err)
	}

	solvePart01(lines)
}

func solvePart01(numbers []string) {
	numCount := len(numbers)
	bitCount := len(numbers[0])
	setBitCount := make([]int, bitCount)

	for _, number := range numbers {
		for bit := 0; bit < bitCount; bit++ {
			if number[bitCount-(bit+1)] == '1' {
				setBitCount[bit]++
			}
		}
	}

	var gamma int
	var epsilon int

	for bit := 0; bit < bitCount; bit++ {
		if setBitCount[bit] > numCount/2 {
			gamma |= (1 << bit)
		} else {
			epsilon |= (1 << bit)
		}
	}

	fmt.Println(gamma * epsilon)
}
