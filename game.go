package main

import (
	"math/rand"
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
	food       []Point
}

func SnakeGame() *Game {
	snakeCanvas, _ := sdc.NewCanvas(sdc.Point{1, 1}, sdc.Point{20, 30})
	dataCanvas, _ := sdc.NewCanvas(sdc.Point{1, 35}, sdc.Point{10, 20})
	return &Game{
		snakeField: snakeCanvas,
		dataField:  dataCanvas,
		snake:      NewSnake(Point{5, 5}),
		speed:      1000,
		isOver:     false,
	}
}
func (gm *Game) addRandomFood() {
	for {
		p := Point{rand.Intn(gm.snakeField.Size().Column-1) + 1, rand.Intn(gm.snakeField.Size().Line-1) + 1}
		if !gm.snake.IsSnakePoint(p) {
			gm.food = append(gm.food, p)
			break
		}
	}
}

func (gm *Game) drawBoxes() {
	gm.snakeField.DrawBoxWithTitle("SNAKE GAME")
	gm.dataField.DrawBoxWithTitle("SCORE")
}

func (gm *Game) repaint() {
	var repaintTimer *time.Timer

	resetRepaintTimer := func() {
		repaintTimer = time.NewTimer(time.Duration(gm.speed) * time.Millisecond)
	}
	resetRepaintTimer()

	for {
		gm.snakeField.ClearInner()
		gm.dataField.ClearInner()

		repaintSnake(gm.snakeField, gm.snake)
		repaintFood(gm.snakeField, gm.food)
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
	cnv.DrawPath("*", convertPoints([]Point{snake.Head()}))
	cnv.DrawPath("0", convertPoints(snake.Body()))
}

func repaintFood(cnv sdc.Canvas, food []Point) {
	cnv.DrawPath("$", convertPoints(food))
}

func repaintScore(cnv sdc.Canvas, snakeLen, speed int) {
	cnv.DrawText("Len  : "+strconv.Itoa(snakeLen), sdc.Point{2, 2})
	cnv.DrawText("Speed: "+strconv.Itoa(speed), sdc.Point{3, 2})
}

func (gm *Game) moveSnake(gameOverChanel chan<- bool) {

}

func (gm *Game) generateFood() {
	var foodTimer *time.Timer

	resetFoodTimer := func() {
		foodTimer = time.NewTimer(2 * time.Duration(gm.speed) * time.Millisecond)
	}
	resetFoodTimer()

	for {
		gm.addRandomFood()

		<-foodTimer.C
		if gm.isOver {
			break
		}
		resetFoodTimer()
	}
}

func (gm *Game) printGameOver() {
	gm.dataField.DrawText("GAME OVER", sdc.Point{4, 2})
}

func (gm *Game) Start() {

	var gameOverChanel chan bool = make(chan bool)
	gm.drawBoxes()

	go gm.moveSnake(gameOverChanel)
	go gm.generateFood()
	go gm.repaint()

	<-gameOverChanel

	time.Sleep(1 * time.Second)
	gm.printGameOver()

}
