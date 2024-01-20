package main

import (
	"time"

	"fyne.io/fyne/v2"
)

type Game struct {
	state [][]*Cell
	worldSize int
	play chan bool
}

func (g *Game) init() {
	state := make([][]*Cell, g.worldSize)

	for row := 0; row < g.worldSize; row++ {
		state[row] = make([]*Cell, 0, g.worldSize)
		for col := 0; col < g.worldSize; col++ {
			state[row] = append(state[row], NewCell(
				Pos{
					row: row,
					col: col,
				},
			))
		}
	}
	g.state = state
}

func (g *Game) SerializeState() []fyne.CanvasObject {
	serialized := []fyne.CanvasObject{}
	for row := 0; row < g.worldSize; row++ {
		for col := 0; col < g.worldSize; col ++ {
			serialized = append(serialized, g.state[row][col])

		}
	}
	return serialized
}

func (g *Game) Forward() {
	nextLiveCellPos := []Pos{}
	for i, row := range g.state {
		for j, cell := range row {
			liveCellCount := 0
			for _, r := range []int{i-1, i, i+1} {
				if (r < 0 || r == g.worldSize) {
					continue
				}
				for _, c := range []int{j-1, j ,j+1} {
					if (c < 0 || c == g.worldSize) {
						continue
					}
					if (r == i && c == j) {
						continue
					}
					neighborCell := g.state[r][c]
					if neighborCell.Live {
						liveCellCount++
					}
				}
			}
			if cell.Live && (liveCellCount == 2 || liveCellCount == 3) {
				nextLiveCellPos = append(nextLiveCellPos, Pos{i, j})
			} else {
				if (liveCellCount == 3) {
					nextLiveCellPos = append(nextLiveCellPos, Pos{i, j})
				}
			}
		}
	}
	for _, row := range g.state {
		for _, cell := range row {
			cell.Die()
		}
	}

	for _, pos := range nextLiveCellPos {
		g.state[pos.row][pos.col].Revive()
	}
}

func (g *Game) Render() {
	for _, row := range g.state {
		for _, cell := range row {
			cell.Update()
		}
	}
}

func (g *Game) Play() {
	ticker := time.NewTicker(time.Second)
	g.play = make(chan bool)

	for {
		select {
		case <- g.play:
			ticker.Stop()
			return
		case <- ticker.C:
			g.Forward()
			g.Render()
		}
	}
}