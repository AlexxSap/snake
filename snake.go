package main

type Point struct {
	X, Y uint8
}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Snake struct {
	body []Point
}

func NewSnake(startPoint Point) Snake {
	return Snake{body: []Point{startPoint}}
}

func (snk Snake) Eat(food Point) {
	snk.body = append([]Point{food}, snk.body...)
}

func (snk Snake) Head() Point {
	return snk.body[0]
}
func (snk Snake) Body() []Point {
	return snk.body[1:]
}

func (snk Snake) Len() int {
	return len(snk.body)
}

func (snk Snake) Move(dir Direction) {
	prev := snk.Head()
	snk.body[0] = func(dir Direction, prev Point) Point {
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
	}(dir, prev)

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
