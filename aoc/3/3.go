package aoc

import (
	"fmt"
	"khhk-adventofcode2024/reader"
	"regexp"
	"strconv"
)

func Aoc() {
	path := "aoc/3"
	// lines := reader.AllText(path + "/example2.txt")
	lines := reader.AllText(path + "/input.txt")

	sum := part1(lines)

	fmt.Println("Sum 1: ", sum)

	sum2 := part2(lines)
	fmt.Println("Sum 2: ", sum2)
}

func part1(text string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(text, -1)

	sum := 0
	for _, match := range matches {
		sum += mul(match)
	}

	return sum
}

func part2(text string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := r.FindAllStringSubmatch(text, -1)

	sum := 0
	do := true
	for _, match := range matches {
		if match[0] == "do()" {
			do = true
		} else if match[0] == "don't()" {
			do = false
		} else if do {
			sum += mul(match)
		}
	}

	return sum
}

func mul(match []string) int {
	a, _ := strconv.Atoi(match[1])
	b, _ := strconv.Atoi(match[2])
	return a * b
}
