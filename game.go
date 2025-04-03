package main

import "github.com/hajimehoshi/ebiten/v2"

type tile struct {
	x             int
	y             int
	isMine        bool
	isUncovered   bool
	isFlagged     bool
	neighborMines int
	tileImage     int
}

func (t *tile) updateTile(x, y int) {
	t.x = x
	t.y = y
}

func newTile() tile {
	return tile{
		x:             -1,
		y:             -1,
		isMine:        false,
		neighborMines: -1,
		tileImage:     0,
	}
}

type grid struct {
	tiles    []tile
	offset   int
	tileSize int
	gridSize int
}

func newGrid() grid {
	return grid{
		tiles:    []tile{},
		offset:   16,
		tileSize: 16,
		gridSize: 144,
	}
}

func (gr *grid) populateGrid() {
	t := newTile()
	for x := gr.offset; x <= gr.gridSize; x += gr.tileSize {
		for y := gr.offset; y <= gr.gridSize; y += gr.tileSize {
			t.updateTile((x-gr.offset)/gr.tileSize, (y-gr.offset)/gr.tileSize)
			gr.tiles = append(gr.tiles, t)
		}
	}
}

func (gr *grid) checkGrid(in input) grid {
	for t := range gr.tiles {
		if gr.tiles[t].x == in.tileClick[0] && gr.tiles[t].y == in.tileClick[1] {
			gr.tiles[t].isUncovered = true
			if in.button == ebiten.MouseButtonLeft {
				gr.tiles[t].tileImage = 1
			}
			if in.button == ebiten.MouseButtonRight {
				gr.tiles[t].tileImage = 2
			}
		}
	}
	return *gr
}
