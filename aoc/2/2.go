package aoc

import (
	"fmt"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"slices"
	"strconv"
	"strings"
)

func Aoc() {
	lines := reader.Lines("aoc/2/input.txt")

	sum := 0
	for _, line := range lines {
		values := util.Map(strings.Split(line, " "), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		sum += Line(values)
	}

	fmt.Println("Sum 1: ", sum)

	sum2 := 0

	for _, line := range lines {
		values := util.Map(strings.Split(line, " "), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		sum2 += Line2(values)
	}

	fmt.Println("Sum 2: ", sum2)
}

func Line(values []int) int {
	trend := 0
	useTrend := false
	for i := 1; i < len(values); i++ {
		prev := values[i-1]
		cur := values[i]
		diff := prev - cur
		if diff == 0 || util.Abs(diff) > 3 {
			return 0
		}

		nextTrend := Trend(diff)
		if useTrend && nextTrend != trend {
			return 0
		}

		trend = nextTrend
		useTrend = true
	}

	return 1
}

func Line2(values []int) int {
	result1 := Line(values)
	if result1 == 1 {
		return 1
	}

	for i := 0; i < len(values); i++ {
		fixedValues := slices.Concat(values[:i], values[i+1:])
		result2 := Line(fixedValues)
		if result2 == 1 {
			return 1
		}
	}

	return 0
}

func Trend(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	panic(n)
}
