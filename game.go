package main

import (
	"math/rand/v2"
)

type tile struct {
	id            int
	x             int
	y             int
	coord         [2]int
	isMine        bool
	isUncovered   bool
	isFlagged     bool
	neighborMines int
	tileImage     int
}

func (t *tile) updateTile(x, y, id int) {
	t.x = x
	t.y = y
	t.id = id
	t.coord = [2]int{x, y}
}

func newTile() tile {
	return tile{
		id:            -1,
		x:             -1,
		y:             -1,
		isMine:        false,
		isUncovered:   false,
		isFlagged:     false,
		neighborMines: 4,
		tileImage:     13,
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
		offset:   32,
		tileSize: 32,
		gridSize: 288,
	}
}

func (gr *grid) populateGrid() {
	t := newTile()
	id := -1
	for x := gr.offset; x <= gr.gridSize; x += gr.tileSize {
		for y := gr.offset; y <= gr.gridSize; y += gr.tileSize {
			id += 1
			t.updateTile((x-gr.offset)/gr.tileSize, (y-gr.offset)/gr.tileSize, id)
			gr.tiles = append(gr.tiles, t)
		}
	}
	gr.generateMines()
	gr.neighborNumbers()
}

func (gr *grid) generateMines() {
	mineLocations := randomNumbers(80, 10)
	for t := range gr.tiles {
		for i := range mineLocations {
			if gr.tiles[t].id == mineLocations[i] {
				gr.tiles[t].isMine = true
			}
		}
	}
}

func (gr *grid) neighborNumbers() {
	var mineCounter int
	neighborsId := []int{-10, -1, 8, -9, 1, -8, 9, 10}
	neighborCoord := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 1}, {-1, 1}, {1, 0}, {1, 1}}
	for t := range gr.tiles {
		for i := range neighborsId {
			nbTile := []int{gr.tiles[t].x + neighborCoord[i][0], gr.tiles[t].y + neighborCoord[i][1]}
			xBounds := nbTile[0] >= 0 && nbTile[0] <= 8
			yBounds := nbTile[1] >= 0 && nbTile[1] <= 8
			if xBounds && yBounds {
				if gr.tiles[t+neighborsId[i]].isMine {
					mineCounter += 1
				}
			}
		}
		gr.tiles[t].neighborMines = mineCounter
		mineCounter = 0
	}
}
func randomNumbers(maxNum, count int) []int {
	selections := make([]int, 0, count)
	numberPool := make([]int, maxNum)
	for i := range numberPool {
		numberPool[i] = i
	}
	for i := 0; i < count; i++ {
		idx := rand.IntN(len(numberPool))
		value := numberPool[idx]
		selections = append(selections, value)

		numberPool[idx] = numberPool[len(numberPool)-1]
		numberPool = numberPool[:len(numberPool)-1]
	}
	return selections
}

func (gr *grid) checkGrid(in input) grid {
	for t := range gr.tiles {
		if gr.tiles[t].x == in.tileClick[0] && gr.tiles[t].y == in.tileClick[1] {
			if in.mouseButtonLeft {
				gr.identifyTileClicked(t)
			}
			if in.mouseButtonRight {
				gr.flag(t)
			}
		}
	}
	return *gr
}

func (gr *grid) identifyTileClicked(t int) {
	if gr.tiles[t].isMine {
		gr.revealMines(t)
		gr.tiles[t].tileImage = 11
	}
	if !gr.tiles[t].isUncovered && !gr.tiles[t].isMine {
		gr.tiles[t].tileImage = gr.tiles[t].neighborMines
		gr.tiles[t].isUncovered = true
	}
}

func (gr *grid) flag(t int) {
	if gr.tiles[t].tileImage == 9 {
		gr.tiles[t].tileImage = 13
		gr.tiles[t].isFlagged = false
	} else if gr.tiles[t].tileImage == 13 {
		gr.tiles[t].tileImage = 9
		gr.tiles[t].isFlagged = true
	} else {
		return
	}
}

func (gr *grid) revealMines(tClick int) {
	for t := range gr.tiles {
		if gr.tiles[t].isMine && t != tClick && !gr.tiles[t].isFlagged {
			gr.tiles[t].tileImage = 10
		}
	}
}
