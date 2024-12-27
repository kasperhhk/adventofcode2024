package fields

type Direction struct {
	Drow int
	Dcol int
}

var UP = Direction{-1, 0}
var DOWN = Direction{1, 0}
var LEFT = Direction{0, -1}
var RIGHT = Direction{0, 1}
