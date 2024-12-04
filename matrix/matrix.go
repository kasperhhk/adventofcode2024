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
