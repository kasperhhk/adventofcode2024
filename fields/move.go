package fields

func Axis4[T any](field [][]T, pos Position) (positions []Position) {
	dirs := []Direction{UP, DOWN, LEFT, RIGHT}
	for _, dir := range dirs {
		p := Move(pos, dir)
		if Within(field, p) {
			// fmt.Println("Axis4 appending", p.Row, p.Col)
			positions = append(positions, p)
		}
	}

	return
}
