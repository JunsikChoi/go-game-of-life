package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	worldSize := 30
	myApp := app.New()
	myWindow := myApp.NewWindow("Conway's Game of Life")
	
	game := Game{worldSize: worldSize}
	game.init()

	cells := game.SerializeState()

	grid := container.New(layout.NewGridLayout(worldSize), cells...)

	forwardBtn := widget.NewButton("Play", func() {
		go game.Play()
	})

	stopBtn := widget.NewButton("Stop", func() {
		game.play <- true
	})

	content := container.New(layout.NewVBoxLayout(), grid, forwardBtn, stopBtn)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

type Pos struct {
	row int
	col int
}
