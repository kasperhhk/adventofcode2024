package aoc

import (
	"fmt"
	"khhk-adventofcode2024/matrix"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"slices"
	"strconv"
	"strings"
)

func Aoc1() {
	lines := reader.Lines("aoc/1/input.txt")

	pairs := util.Map(lines, func(line string) []int {
		split := strings.Split(line, " ")

		filtered := util.Filter(split, func(cell string) bool {
			return len(strings.TrimSpace(cell)) > 0
		})

		pair := util.Map(filtered, func(cell string) int {
			i, _ := strconv.Atoi(cell)
			return i
		})
		return pair
	})

	// matrix.PrintMatrix(pairs)

	transposed := matrix.Transpose(pairs, len(pairs), len(pairs[0]))
	// matrix.PrintMatrix(transposed)

	slices.Sort(transposed[0])
	slices.Sort(transposed[1])

	l := transposed[0]
	r := transposed[1]
	sum := 0
	for i := 0; i < len(transposed[0]); i++ {
		diff := l[i] - r[i]
		sum += max(diff, -diff)
	}

	fmt.Println("Sum: ", sum)

	fmt.Println("--------------- PART 2 ---------------")

	fmap := make(map[int]int)
	sum2 := 0

	for _, v := range l {
		freq, ok := fmap[v]

		if !ok {
			for _, vr := range r {
				if v == vr {
					freq++
				}
			}
			fmap[v] = freq
		}
		sum2 += freq * v
	}

	fmt.Println("Sum2: ", sum2)
}
