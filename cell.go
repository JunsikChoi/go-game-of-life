package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Cell struct {
	widget.BaseWidget
	rect *canvas.Rectangle
	Live bool
	Pos Pos
}

func NewCell(pos Pos) *Cell {
	rect := canvas.NewRectangle(color.Black)
	rect.SetMinSize(fyne.NewSize(20,20))
	rect.StrokeWidth = 1
	rect.StrokeColor = color.White
	cell := &Cell{
		Live: false,
		rect: rect,
		Pos: pos,
	}
	
	cell.ExtendBaseWidget(cell)
	return cell
}

func (c *Cell) CreateRenderer() fyne.WidgetRenderer {
	if c.rect == nil {
		return nil
	}

	return widget.NewSimpleRenderer(c.rect)
}

func (c *Cell) Tapped(*fyne.PointEvent) {
	if (c.Live) {
		c.Die()
	} else {
		c.Revive()
	}
	c.Update()
}

func (c *Cell) Die() {
	c.Live = false
}

func (c *Cell) Revive() {
	c.Live = true
}

func (c *Cell) Update() {
	if (c.Live) {
		c.rect.FillColor = color.White
	} else {
		c.rect.FillColor = color.Black
	}
	c.rect.Refresh()
}