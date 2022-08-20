package snake

type Point struct {
	X, Y uint8
}

type Snake struct {
	body []Point
}

func NewSnake(startPoint Point) Snake {
	return Snake{body: []Point{startPoint}}
}

func (snk Snake) Eat(food Point) {
	snk.body = append([]Point{food}, snk.body...)
}

func (snk Snake) Len() int {
	return len(snk.body)
}
