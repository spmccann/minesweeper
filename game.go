package main

import "fmt"

type tile struct {
	x             int
	y             int
	isMine        bool
	isCovered     bool
	isFlagged     bool
	neighborMines int
}

func (t *tile) updateTile(x, y int) {
	t.x = x
	t.y = y
}

type grid struct {
	tiles []tile
}

func (g *grid) populateGrid() {
	t := tile{}
	current := 16
	maxSize := 144
	for x := current; x <= maxSize; x += current {
		for y := current; y <= maxSize; y += current {
			t.updateTile((x-16)/16, (y-16)/16)
			g.tiles = append(g.tiles, t)
		}
	}
}

func (g *grid) checkGrid() (int, int) {
	xLoc, yLoc := calcTileClicked()
	for t := range g.tiles {
		if g.tiles[t].x == xLoc && g.tiles[t].y == yLoc {
			fmt.Println("clicked the tile", xLoc, yLoc)
			return xLoc, yLoc
		}
	}
	return -1, -1
}

func calcTileClicked() (int, int) {
	xLoc := (locationClicked[0] - 16) / 16
	yLoc := (locationClicked[1] - 16) / 16
	return xLoc, yLoc
}

func updateTileImage() (int, int) {
	g := grid{}
	g.populateGrid()
	return g.checkGrid()
}
