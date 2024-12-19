package aoc

import (
	"fmt"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"strconv"
	"strings"
)

func Aoc() {
	path := "aoc/7"
	// lines := reader.Lines(path + "/example.txt")
	lines := reader.Lines(path + "/input.txt")

	equations := parse(lines)

	sum := part1(equations)
	fmt.Println("Sum 1: ", sum)

	sum2 := part2(equations)
	fmt.Println("Sum 2: ", sum2)
}

/* PART 1 SPECIFICS */

func part1(equations []Equation) int {
	sum := 0

	for _, eq := range equations {
		if check(eq) {
			sum += eq.target
		}
	}

	return sum
}

func check(equation Equation) bool {
	return recursiveCheck(equation.target, equation.values[0], equation.values[1:])
}

func recursiveCheck(target int, current int, remainder []int) bool {
	if current > target {
		return false
	}

	if len(remainder) == 0 {
		return current == target
	}

	nextMult := current * remainder[0]
	nextAdd := current + remainder[0]
	nextRemainder := remainder[1:]

	return recursiveCheck(target, nextMult, nextRemainder) || recursiveCheck(target, nextAdd, nextRemainder)
}

/* PART 2 SPECIFICS */

func part2(equations []Equation) int {
	sum := 0

	for _, eq := range equations {
		if check2(eq) {
			sum += eq.target
		}
	}

	return sum
}

func check2(equation Equation) bool {
	return recursiveCheck2(equation.target, equation.values[0], equation.values[1:])
}

func recursiveCheck2(target int, current int, remainder []int) bool {
	// fmt.Println("Target: ", target, ", Current: ", current, ", Remainder: ", remainder)
	if current > target {
		return false
	}

	if len(remainder) == 0 {
		return current == target
	}

	nextConcat, _ := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(remainder[0]))
	nextMult := current * remainder[0]
	nextAdd := current + remainder[0]

	nextRemainder := remainder[1:]

	return recursiveCheck2(target, nextConcat, nextRemainder) || recursiveCheck2(target, nextMult, nextRemainder) || recursiveCheck2(target, nextAdd, nextRemainder)
}

/* PARSING AND MODELLING */

func parse(lines []string) (equations []Equation) {
	for _, line := range lines {
		equations = append(equations, parseEquation(line))
	}

	return
}

func parseEquation(line string) Equation {
	spl1 := strings.Split(line, ":")
	left, rest := spl1[0], strings.TrimSpace(spl1[1])

	target, _ := strconv.Atoi(left)

	nums := util.Map(strings.Split(rest, " "), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	return Equation{target, nums}
}

type Equation struct {
	target int
	values []int
}
