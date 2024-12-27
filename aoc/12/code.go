package aoc

import (
	"fmt"
	"khhk-adventofcode2024/fields"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"strings"
)

func Aoc() {
	path := "aoc/12"
	// lines := reader.Lines(path + "/example.txt")
	lines := reader.Lines(path + "/input.txt")

	field := parse(lines)

	//matrix.PrintMatrix(arr)

	sum := part1(field)
	fmt.Println("Sum 1: ", sum)

	sum2 := part2()
	fmt.Println("Sum 2: ", sum2)
}

/* PART 1 SPECIFICS */

func part1(field [][]string) int {
	sum := 0

	rs := regions(field)
	for _, region := range rs {
		sum += area(region) * perimiter(field, region)
	}

	return sum
}

func area(region Region) int {
	return len(region.positions)
}

func perimiter(field [][]string, region Region) int {
	sum := 0

	for _, p := range region.positions {
		sum += perimiterSingle(field, p)
	}

	return sum
}

func perimiterSingle(field [][]string, pos fields.Position) int {
	val := 4

	neighbours := fields.Axis4(field, pos)
	for _, n := range neighbours {
		if field[n.Row][n.Col] == field[pos.Row][pos.Col] {
			val--
		}
	}

	return val
}

/* PART 2 SPECIFICS */

func part2() int {
	sum := 0

	return sum
}

/* MODELLING */
func regions(field [][]string) (regions []Region) {
	seen := make(map[fields.Position]bool)

	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[0]); c++ {
			pos := fields.Position{Row: r, Col: c}
			_, ok := seen[pos]
			if ok {
				continue
			}

			region := mapRegion(field, pos, seen)
			regions = append(regions, region)
		}
	}

	return
}

func mapRegion(field [][]string, start fields.Position, seen map[fields.Position]bool) Region {
	positions := explore(field, start, seen)
	return Region{positions}
}

func explore(field [][]string, pos fields.Position, seen map[fields.Position]bool) (output []fields.Position) {
	_, ok := seen[pos]
	if ok {
		return
	}

	seen[pos] = true
	output = append(output, pos)

	neighbours := fields.Axis4(field, pos)
	for _, n := range neighbours {
		if field[n.Row][n.Col] == field[pos.Row][pos.Col] {
			res := explore(field, n, seen)
			output = append(output, res...)
		}
	}

	return
}

type Region struct {
	positions []fields.Position
}

/* PARSING */

func parse(lines []string) [][]string {
	return util.Map(lines, func(s string) []string {
		return strings.Split(s, "")
	})
}
