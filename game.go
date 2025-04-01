package main

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

type grid struct {
	tiles []tile
}

func (gr *grid) populateGrid() {
	t := tile{}
	current := 16
	maxSize := 144
	for x := current; x <= maxSize; x += current {
		for y := current; y <= maxSize; y += current {
			t.updateTile((x-16)/16, (y-16)/16)
			gr.tiles = append(gr.tiles, t)
		}
	}
}

func (gr *grid) checkGrid() grid {
	xLoc, yLoc := calcTileClicked()
	for t := range gr.tiles {
		if gr.tiles[t].x == xLoc && gr.tiles[t].y == yLoc {
			gr.tiles[t].isUncovered = true
			gr.tiles[t].tileImage = 1
		}
	}
	return *gr
}

func calcTileClicked() (int, int) {
	xLoc := (locationClicked[0] - 16) / 16
	yLoc := (locationClicked[1] - 16) / 16
	return xLoc, yLoc
}

func updateTileData(newGame bool, gr grid) grid {
	if newGame {
		gr := grid{}
		gr.populateGrid()
		return gr.checkGrid()
	}
	return gr.checkGrid()
}
