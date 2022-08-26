package main

import (
	"strconv"
	"time"

	sdc "github.com/AlexxSap/SiDCo"
)

func (gm *Game) repaint() {
	var repaintTimer *time.Timer

	resetRepaintTimer := func() {
		repaintTimer = time.NewTimer(time.Duration(gm.speed) * time.Millisecond)
	}
	resetRepaintTimer()

	prevSpeed := 0
	for {
		gm.snakeField.ClearInner()

		repaintSnake(gm.snakeField, gm.snake)
		repaintFood(gm.snakeField, gm.food)
		if prevSpeed != gm.speed {
			gm.dataField.ClearInner()
			repaintScore(gm.dataField, gm.snake.Len(), gm.speed)
			prevSpeed = gm.speed
		}

		if gm.isOver {
			break
		}

		<-repaintTimer.C
		resetRepaintTimer()
	}
}

func (gm *Game) drawBoxes() {
	gm.snakeField.DrawBoxWithTitle("SNAKE GAME")
	gm.dataField.DrawBoxWithTitle("SCORE")
}

func repaintSnake(cnv sdc.Canvas, snake *Snake) {
	cnv.DrawPath("o", convertPoints(snake.Points()))
}

func repaintFood(cnv sdc.Canvas, food []Point) {
	cnv.DrawPath("$", convertPoints(food))
}

func repaintScore(cnv sdc.Canvas, snakeLen, speed int) {
	cnv.DrawText("Len  : "+strconv.Itoa(snakeLen), sdc.Point{Line: 2, Column: 2})
	cnv.DrawText("Speed: "+strconv.Itoa(1000-speed), sdc.Point{Line: 3, Column: 2})
}

func (gm *Game) printGameOver() {
	gm.dataField.DrawText("<===GAME OVER===>", sdc.Point{Line: 5, Column: 2})
}
