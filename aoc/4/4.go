package aoc

import (
	"fmt"
	"khhk-adventofcode2024/matrix"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"strings"
)

func Aoc() {
	path := "aoc/4"
	// lines := reader.Lines(path + "/example.txt")
	lines := reader.Lines(path + "/input.txt")

	sum := part1(lines)
	fmt.Println("Sum 1: ", sum)

	sum2 := part2(lines)
	fmt.Println("Sum 2: ", sum2)
}

func part2(lines []string) int {
	mapped := util.Map(lines, func(s string) []string {
		return strings.Split(s, "")
	})

	sum := 0
	for r, row := range mapped {
		for c, value := range row {
			if value == "A" {
				sum += check2(mapped, r, c, len(mapped), len(row))
			}
		}
	}

	return sum
}

func check2(mapped [][]string, r int, c int, rows int, cols int) int {
	if r-1 < 0 || r+1 >= rows || c-1 < 0 || c+1 >= cols {
		return 0
	}

	topleft := mapped[r-1][c-1]
	topright := mapped[r-1][c+1]
	botleft := mapped[r+1][c-1]
	botright := mapped[r+1][c+1]

	diag1 := topleft == "M" && botright == "S" || topleft == "S" && botright == "M"
	diag2 := botleft == "M" && topright == "S" || botleft == "S" && topright == "M"

	if diag1 && diag2 {
		return 1
	}

	return 0
}

func part1(lines []string) int {
	orig := util.Map(lines, func(s string) []string {
		return strings.Split(s, "")
	})

	sum := 0
	mat := orig
	for i := 0; i < 4; i++ {
		sum += check(mat)
		mat = matrix.Rotate(mat, len(mat), len(mat[0]))
	}

	return sum
}

func check(mat [][]string) int {
	sum := 0

	for _, row := range mat {
		joined := strings.Join(row, "")
		sum += strings.Count(joined, "XMAS")
	}

	diagonals := matrix.Diagonals(mat, len(mat), len(mat[0]))
	for _, row := range diagonals {
		joined := strings.Join(row, "")
		sum += strings.Count(joined, "XMAS")
	}

	return sum
}
