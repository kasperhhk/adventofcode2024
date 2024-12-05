package matrix

import "fmt"

func Transpose[A any](source [][]A, rows int, cols int) [][]A {
	outer := make([][]A, cols)
	for c := 0; c < cols; c++ {
		inner := make([]A, rows)

		for r := 0; r < rows; r++ {
			inner[r] = source[r][c]
		}

		outer[c] = inner
	}

	return outer
}

func PrintMatrix[A any](source [][]A) {
	for _, row := range source {
		fmt.Println(row)
	}
}

func Rotate[A any](source [][]A, rows int, cols int) [][]A {
	outer := make([][]A, cols)
	for c := 0; c < cols; c++ {
		inner := make([]A, rows)
		outer[c] = inner
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			outer[c][rows-1-r] = source[r][c]
		}
	}

	return outer
}

func Diagonals[A any](source [][]A, rows int, cols int) [][]A {
	var diagonals [][]A

	for r := 0; r < rows; r++ {
		d := diagonal(source, rows, cols, r, 0)
		diagonals = append(diagonals, d)
	}

	for c := 1; c < cols; c++ {
		d := diagonal(source, rows, cols, rows-1, c)
		diagonals = append(diagonals, d)
	}

	return diagonals
}

func diagonal[A any](source [][]A, _ int, cols int, rstart int, cstart int) []A {
	var diagonal []A

	r := rstart
	c := cstart
	for r >= 0 && c < cols {
		diagonal = append(diagonal, source[r][c])
		r--
		c++
	}

	return diagonal
}
