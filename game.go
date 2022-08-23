package main

import sdc "github.com/AlexxSap/SiDCo"

type Game struct {
	snakeField sdc.Canvas
	dataField  sdc.Canvas
}

func NewGame() Game {
	snake, _ := sdc.NewCanvas(sdc.Point{1, 1}, sdc.Point{20, 30})
	data, _ := sdc.NewCanvas(sdc.Point{1, 35}, sdc.Point{10, 20})
	return Game{
		snakeField: snake,
		dataField:  data,
	}
}

func (gm Game) drawBoxes() {
	gm.snakeField.DrawBoxWithTitle("SNAKE GAME")
	gm.dataField.DrawBoxWithTitle("SCORE")
}

func (gm Game) Start() {
	gm.drawBoxes()
}
