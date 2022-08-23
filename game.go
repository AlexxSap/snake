package main

import (
	sdc "github.com/AlexxSap/SiDCo"
)

type Game struct {
	snakeField sdc.Canvas
	dataField  sdc.Canvas
	snake      Snake
}

func NewGame() Game {
	snake, _ := sdc.NewCanvas(sdc.Point{1, 1}, sdc.Point{20, 30})
	data, _ := sdc.NewCanvas(sdc.Point{1, 35}, sdc.Point{10, 20})
	return Game{
		snakeField: snake,
		dataField:  data,
		snake:      NewSnake(Point{5, 5}),
	}
}

func (gm Game) drawBoxes() {
	gm.snakeField.DrawBoxWithTitle("SNAKE GAME")
	gm.dataField.DrawBoxWithTitle("SCORE")
}

func (gm Game) repaint() {
	repaintSnake(gm.snakeField, gm.snake)
	repaintScore(gm.dataField, gm.snake)
}

func convertPoints(points []Point) []sdc.Point {
	result := make([]sdc.Point, 0, len(points))
	for _, p := range points {
		result = append(result, sdc.Point{Line: int(p.Y), Column: int(p.X)})
	}
	return result
}

func repaintSnake(cnv sdc.Canvas, snake Snake) {
	cnv.ClearInner()
	cnv.DrawPath("^", convertPoints(snake.Points()))
}

func repaintScore(cnv sdc.Canvas, snake Snake) {
	cnv.ClearInner()
}

func (gm Game) Start() {
	gm.drawBoxes()

	gm.repaint()
}
