package main

type Point struct {
	X, Y int
}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

func MovePoint(dir Direction, prev Point) Point {
	switch dir {
	case Up:
		return Point{prev.X, prev.Y - 1}
	case Down:
		return Point{prev.X, prev.Y + 1}
	case Right:
		return Point{prev.X + 1, prev.Y}
	case Left:
		return Point{prev.X - 1, prev.Y}
	}
	return prev
}

type Snake struct {
	body []Point
}

// TODO сделать дэфолтную змейку длиннее
func NewSnake(startPoint Point) Snake {
	return Snake{body: []Point{startPoint}}
}

func (snk Snake) Eat(food Point) {
	snk.body = append([]Point{food}, snk.body...)
}

func (snk Snake) Points() []Point {
	return snk.body
}

func (snk Snake) Head() Point {
	return snk.body[0]
}
func (snk Snake) Body() []Point {
	return snk.body[1:]
}

func (snk Snake) IsSnakePoint(p Point) bool {
	for _, s := range snk.body {
		if s == p {
			return true
		}
	}
	return false
}

func (snk Snake) Len() int {
	return len(snk.body)
}

func (snk Snake) Move(dir Direction) {
	prev := snk.Head()
	snk.body[0] = MovePoint(dir, prev)

	for i := 1; i < snk.Len(); i++ {
		snk.body[i], prev = prev, snk.body[i]
	}
}

func (snk Snake) IsSelfBite() bool {
	h := snk.Head()
	for _, p := range snk.Body() {
		if p == h {
			return true
		}
	}
	return false
}
