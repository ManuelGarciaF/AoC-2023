package utils

type Coord struct {
	X int
	Y int
}

func (c Coord) Add(other Coord) Coord {
	return Coord{c.X + other.X, c.Y + other.Y}
}

func (c Coord) Sub(other Coord) Coord {
	return Coord{c.X - other.X, c.Y - other.Y}
}

func (c Coord) Scale(factor int) Coord {
	return Coord{c.X * factor, c.Y * factor}
}

// Suposes (0,0) at top-left corner, with x going downwards
// and y going to the right
var Offsets = map[Direction]Coord{
	UP:    {X: -1, Y: 0},
	DOWN:  {X: 1, Y: 0},
	LEFT:  {X: 0, Y: -1},
	RIGHT: {X: 0, Y: 1},
}

var DirFromOffset = map[Coord]Direction{
	{X: -1, Y: 0}: UP,
	{X: 1, Y: 0}:  DOWN,
	{X: 0, Y: -1}: LEFT,
	{X: 0, Y: 1}:  RIGHT,
}


type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var InverseDir = map[Direction]Direction{
	UP:    DOWN,
	DOWN:  UP,
	LEFT:  RIGHT,
	RIGHT: LEFT,
}
