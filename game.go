package main

import (
	"time"

	sdc "github.com/AlexxSap/SiDCo"
	"github.com/eiannone/keyboard"
)

type Game struct {
	snakeField sdc.Canvas
	dataField  sdc.Canvas
	snake      *Snake
	speed      int
	isOver     bool
	food       []Point
	direction  Direction
}

func SnakeGame() *Game {
	snakeCanvas, _ := sdc.NewCanvas(sdc.Point{Line: 1, Column: 1}, sdc.Point{Line: 15, Column: 25})
	dataCanvas, _ := sdc.NewCanvas(sdc.Point{Line: 1, Column: 30}, sdc.Point{Line: 10, Column: 20})

	return &Game{
		snakeField: snakeCanvas,
		dataField:  dataCanvas,
		snake:      NewSnake([]Point{Point{5, 5}, Point{4, 5}, Point{3, 5}}),
		speed:      500,
		isOver:     false,
		direction:  Right,
	}
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

func convertPoints(points []Point) []sdc.Point {
	result := make([]sdc.Point, 0, len(points))
	for _, p := range points {
		result = append(result, sdc.Point{Line: int(p.Y), Column: int(p.X)})
	}
	return result
}

func (gm *Game) isSnakeOutOfBox() bool {
	head := gm.snake.Head()
	size := gm.snakeField.Size()
	return head.X < 1 || head.Y < 1 || head.X > size.Column-1 || head.Y > size.Line-1
}

func (gm *Game) isSnakeDead() bool {
	return gm.snake.IsSelfBite() || gm.isSnakeOutOfBox()
}

func (gm *Game) checkKeyPress() {
	for {
		if gm.isOver {
			break
		}

		_, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		prevDir := gm.direction

		switch key {
		case keyboard.KeyArrowLeft:
			if prevDir != Right {
				gm.direction = Left
			}
		case keyboard.KeyArrowUp:
			if prevDir != Down {
				gm.direction = Up
			}
		case keyboard.KeyArrowDown:
			if prevDir != Up {
				gm.direction = Down
			}
		case keyboard.KeyArrowRight:
			if prevDir != Left {
				gm.direction = Right
			}
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
			gm.removeFood(nextPoint)
			gm.speed = gm.speed - 20
		} else {
			gm.snake.Move(gm.direction)
		}

		if gm.isSnakeDead() {
			gameOverChanel <- true
			break
		}

		resetMoveTimer()
	}
}
