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

// Suposes (0,0) at top-left corner, with x going downwards
// and y going to the right
var Offsets = map[Direction]Coord{
	UP:    {X: 0, Y: -1},
	DOWN:  {X: 0, Y: 1},
	LEFT:  {X: -1, Y: 0},
	RIGHT: {X: 1, Y: 0},
}

var DirFromOffset = map[Coord]Direction{
	{X: 0, Y: -1}: UP,
	{X: 0, Y: 1}:  DOWN,
	{X: -1, Y: 0}: LEFT,
	{X: 1, Y: 0}:  RIGHT,
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
