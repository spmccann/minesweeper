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
	tileSize int
	gridSize int
}

func newGrid() grid {
	return grid{
		tiles:    []tile{},
		tileSize: 16,
		gridSize: 144,
	}
}

func (gr *grid) populateGrid() {
	t := newTile()
	for x := gr.tileSize; x <= gr.gridSize; x += gr.tileSize {
		for y := gr.tileSize; y <= gr.gridSize; y += gr.tileSize {
			t.updateTile((x-16)/16, (y-16)/16)
			gr.tiles = append(gr.tiles, t)
		}
	}
}

func (gr *grid) checkGrid(tileClick [2]int) grid {
	for t := range gr.tiles {
		if gr.tiles[t].x == tileClick[0] && gr.tiles[t].y == tileClick[1] {
			gr.tiles[t].isUncovered = true
			gr.tiles[t].tileImage = 1
		}
	}
	return *gr
}
