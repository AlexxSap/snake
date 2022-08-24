package main

import (
	"strconv"
	"time"

	sdc "github.com/AlexxSap/SiDCo"
)

type Game struct {
	snakeField sdc.Canvas
	dataField  sdc.Canvas
	snake      Snake
	speed      int
	isOver     bool
}

func SnakeGame() Game {
	snakeCanvas, _ := sdc.NewCanvas(sdc.Point{1, 1}, sdc.Point{20, 30})
	dataCanvas, _ := sdc.NewCanvas(sdc.Point{1, 35}, sdc.Point{10, 20})
	return Game{
		snakeField: snakeCanvas,
		dataField:  dataCanvas,
		snake:      NewSnake(Point{5, 5}),
		speed:      1000,
		isOver:     false,
	}
}

func (gm Game) drawBoxes() {
	gm.snakeField.DrawBoxWithTitle("SNAKE GAME")
	gm.dataField.DrawBoxWithTitle("SCORE")
}

func (gm Game) repaint() {
	var repaintTimer *time.Timer

	resetRepaintTimer := func() {
		repaintTimer = time.NewTimer(time.Duration(gm.speed) * time.Millisecond)
	}
	resetRepaintTimer()

	for {
		repaintSnake(gm.snakeField, gm.snake)
		repaintScore(gm.dataField, gm.snake.Len(), gm.speed)

		if gm.isOver {
			break
		}

		<-repaintTimer.C
		resetRepaintTimer()
	}
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

func repaintScore(cnv sdc.Canvas, snakeLen, speed int) {
	cnv.ClearInner()
	cnv.DrawText("Len  : "+strconv.Itoa(snakeLen), sdc.Point{2, 2})
	cnv.DrawText("Speed: "+strconv.Itoa(speed), sdc.Point{3, 2})
}

func (gm Game) moveSnake() {

}

func (gm Game) generateFood() {

}

func (gm Game) Start() {
	gm.drawBoxes()

	//go gm.moveSnake()
	//go gm.generateFood()

	gm.repaint()
}
