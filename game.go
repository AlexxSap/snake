package main

import (
	"math/rand"
	"strconv"
	"time"

	sdc "github.com/AlexxSap/SiDCo"
	"github.com/eiannone/keyboard"
)

type Game struct {
	snakeField sdc.Canvas
	dataField  sdc.Canvas
	snake      Snake
	speed      int
	isOver     bool
	food       []Point
	direction  Direction
}

func SnakeGame() *Game {
	snakeCanvas, _ := sdc.NewCanvas(sdc.Point{Line: 1, Column: 1}, sdc.Point{Line: 20, Column: 30})
	dataCanvas, _ := sdc.NewCanvas(sdc.Point{Line: 1, Column: 35}, sdc.Point{Line: 10, Column: 20})

	return &Game{
		snakeField: snakeCanvas,
		dataField:  dataCanvas,
		snake:      NewSnake([]Point{Point{5, 5}, Point{5, 4}, Point{5, 3}}),
		speed:      1000,
		isOver:     false,
		direction:  Right,
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

func (gm *Game) isFood(point Point) bool {
	for _, food := range gm.food {
		if point == food {
			return true
		}
	}
	return false
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

// TODO make snake green
func repaintSnake(cnv sdc.Canvas, snake Snake) {
	//TODO make snake's head painted by diff color?
	cnv.DrawPath("*", convertPoints([]Point{snake.Head()}))
	cnv.DrawPath("0", convertPoints(snake.Body()))
}

// TODO make food yellow
func repaintFood(cnv sdc.Canvas, food []Point) {
	cnv.DrawPath("$", convertPoints(food))
}

func repaintScore(cnv sdc.Canvas, snakeLen, speed int) {
	cnv.DrawText("Len  : "+strconv.Itoa(snakeLen), sdc.Point{Line: 2, Column: 2})
	cnv.DrawText("Speed: "+strconv.Itoa(speed), sdc.Point{Line: 3, Column: 2})
}

func (gm *Game) isSnakeDead() bool {
	// TODO добавить проверку выхода за пределы поля
	return gm.snake.IsSelfBite()
}

func (gm *Game) checkKeyPress() {
	for {
		if gm.isOver {
			break
		}

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowLeft:
			gm.direction = Left
		case keyboard.KeyArrowUp:
			gm.direction = Up
		case keyboard.KeyArrowDown:
			gm.direction = Down
		case keyboard.KeyArrowRight:
			gm.direction = Right
		}
	}
}

func (gm *Game) moveSnake(gameOverChanel chan<- bool) {
	var moveTimer *time.Timer
	resetMoveTimer := func() {
		moveTimer = time.NewTimer(time.Duration(gm.speed) * time.Millisecond)
	}
	resetMoveTimer()

	for {
		<-moveTimer.C

		nextPoint := MovePoint(gm.direction, gm.snake.Head())
		if gm.isFood(nextPoint) {
			gm.snake.Eat(nextPoint)
		} else {
			gm.snake.Move(gm.direction)
		}

		if gm.isSnakeDead() {
			gameOverChanel <- true
			break
		}
	}
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
	gm.dataField.DrawText("<===GAME OVER===>", sdc.Point{Line: 5, Column: 2})
}

func (gm *Game) Start() {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var gameOverChanel chan bool = make(chan bool)
	gm.drawBoxes()

	go gm.checkKeyPress()
	go gm.moveSnake(gameOverChanel)
	go gm.generateFood()
	go gm.repaint()

	<-gameOverChanel
	gm.isOver = true

	time.Sleep(1 * time.Second)
	gm.printGameOver()

}
