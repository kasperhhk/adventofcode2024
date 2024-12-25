package aoc

import (
	"fmt"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"strconv"
	"strings"
)

func Aoc() {
	path := "aoc/10"
	// lines := reader.Lines(path + "/example.txt")
	// lines := reader.Lines(path + "/example2.txt")
	lines := reader.Lines(path + "/input.txt")

	arr := parse(lines)
	//matrix.PrintMatrix(arr)

	trails := processTrails(arr)
	//printTrails(trails)

	sum := part1(arr, trails)
	fmt.Println("Sum 1: ", sum)

	sum2 := part2(arr, trails)
	fmt.Println("Sum 2: ", sum2)
}

/* PART 1 SPECIFICS */

func part1(arr [][]int, trails map[Position]map[Position]int) int {
	sum := 0

	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr); c++ {
			if arr[r][c] != 0 {
				continue
			}

			map9 := trails[Position{r, c}]
			sum += len(map9)
		}
	}

	return sum
}

/* PART 2 SPECIFICS */

func part2(arr [][]int, trails map[Position]map[Position]int) int {
	sum := 0

	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr); c++ {
			if arr[r][c] != 0 {
				continue
			}

			map9 := trails[Position{r, c}]
			for _, n9 := range map9 {
				sum += n9
			}
		}
	}

	return sum
}

/* MODELLING */

func printTrails(trails map[Position]map[Position]int) {
	for pos, map9 := range trails {
		fmt.Println("cell: ", pos)
		for pos9, n9 := range map9 {
			fmt.Println("found ", pos9, n9, " times")
		}
	}
}

func processTrails(arr [][]int) map[Position]map[Position]int {
	trails := make(map[Position]map[Position]int)

	for r := 0; r < len(arr); r++ {
		for c := 0; c < len(arr[0]); c++ {
			processTrail(arr, Position{r, c}, trails)
		}
	}

	return trails
}

func processTrail(arr [][]int, pos Position, trails map[Position]map[Position]int) {
	// check if we've already processed this cell
	_, ok := trails[pos]
	if ok {
		return
	}

	cellMap := make(map[Position]int)
	trails[pos] = cellMap

	if arr[pos.r][pos.c] == 9 {
		cellMap[pos] = 1
	} else {
		ps := paths(arr, pos)
		for _, p := range ps {
			nPos := Position{pos.r + p.dr, pos.c + p.dc}
			processTrail(arr, nPos, trails)
			for k, v := range trails[nPos] {
				me := cellMap[k]
				cellMap[k] = me + v
			}
		}
	}
}

func paths(arr [][]int, pos Position) (directions []Direction) {
	if canPath(arr, pos, UP) {
		directions = append(directions, UP)
	}
	if canPath(arr, pos, DOWN) {
		directions = append(directions, DOWN)
	}
	if canPath(arr, pos, LEFT) {
		directions = append(directions, LEFT)
	}
	if canPath(arr, pos, RIGHT) {
		directions = append(directions, RIGHT)
	}

	return
}

func canPath(arr [][]int, pos Position, direction Direction) bool {
	pos2 := Position{pos.r + direction.dr, pos.c + direction.dc}
	if pos2.r < 0 || pos2.c < 0 || pos2.r >= len(arr) || pos2.c >= len(arr[0]) {
		return false
	}

	return arr[pos.r][pos.c]+1 == arr[pos2.r][pos2.c]
}

type Position struct {
	r int
	c int
}

type Direction struct {
	dr int
	dc int
}

var (
	UP    = Direction{-1, 0}
	DOWN  = Direction{1, 0}
	LEFT  = Direction{0, -1}
	RIGHT = Direction{0, 1}
)

/* PARSING */

func parse(lines []string) [][]int {
	var arr [][]int
	for _, line := range lines {
		arr = append(arr, parseLine(line))
	}

	return arr
}

func parseLine(line string) []int {
	return util.Map(strings.Split(line, ""), func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
}
