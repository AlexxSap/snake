package main

import (
	"math/rand"
	"time"
)

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

func (gm *Game) removeFood(point Point) {
	for i, f := range gm.food {
		if point == f {
			newFood := gm.food[:i]
			newFood = append(newFood, gm.food[i+1:]...)
			gm.food = newFood
			break
		}
	}
}

func (gm *Game) generateFood() {
	var foodTimer *time.Timer

	resetFoodTimer := func() {
		foodTimer = time.NewTimer(15 * time.Duration(gm.speed) * time.Millisecond)
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
