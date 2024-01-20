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
		if !game.isPlay {
			go game.Play()
		}
	})

	stopBtn := widget.NewButton("Stop", func() {
		game.Stop()
	})

	resetBtn := widget.NewButton("Reset", func() {
		game.Reset()
	})

	btnContainer := container.New(layout.NewHBoxLayout(), forwardBtn, stopBtn, resetBtn)
	content := container.New(layout.NewVBoxLayout(), grid, btnContainer)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

type Pos struct {
	row int
	col int
}
