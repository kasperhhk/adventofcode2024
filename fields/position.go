package fields

func Move(pos Position, dir Direction) Position {
	return Position{pos.Row + dir.Drow, pos.Col + dir.Dcol}
}

func Within[T any](field [][]T, pos Position) bool {
	// fmt.Println("within", pos.Row, pos.Col)
	return pos.Row >= 0 && pos.Col >= 0 && pos.Row < len(field) && pos.Col < len(field[0])
}

type Position struct {
	Row int
	Col int
}
