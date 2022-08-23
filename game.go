package main

import (
	"time"

	sdc "github.com/AlexxSap/SiDCo"
)

type Game struct {
	snakeField sdc.Canvas
	dataField  sdc.Canvas
	snake      Snake
	speed      int
}

func NewGame() Game {
	snake, _ := sdc.NewCanvas(sdc.Point{1, 1}, sdc.Point{20, 30})
	data, _ := sdc.NewCanvas(sdc.Point{1, 35}, sdc.Point{10, 20})
	return Game{
		snakeField: snake,
		dataField:  data,
		snake:      NewSnake(Point{5, 5}),
		speed:      500,
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
	cnv.DrawPath("*", convertPoints([]Point{snake.Head()}))
	cnv.DrawPath("0", convertPoints(snake.Body()))
}

func repaintScore(cnv sdc.Canvas, snake Snake) {
	cnv.ClearInner()
	cnv.DrawText("Len  : nnn", sdc.Point{2, 2})
	cnv.DrawText("Speed: nnn", sdc.Point{3, 2})
}

func (gm Game) Start() {
	gm.drawBoxes()

	var globalTimer *time.Timer

	reserGlobalTimer := func() {
		globalTimer = time.NewTimer(time.Duration(gm.speed) * time.Microsecond)
	}
	reserGlobalTimer()

	for {
		<-globalTimer.C

		gm.repaint()
	}

}
