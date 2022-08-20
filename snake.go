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
