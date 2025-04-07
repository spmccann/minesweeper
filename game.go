package main

import (
	"fmt"
	"math/rand/v2"
)

type tile struct {
	id            int
	x             int
	y             int
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
		offset:   16,
		tileSize: 16,
		gridSize: 144,
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
	fmt.Println(gr.tiles)
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
	neighbors := []int{-10, -9, -8, -1, 1, 8, 9, 10}
	for t := range gr.tiles {
		for i := range neighbors {
			if t+neighbors[i] < 81 && t+neighbors[i] > -1 {
				if gr.tiles[t+neighbors[i]].isMine {
					mineCounter += 1
					fmt.Println(t, "found mine at", t+neighbors[i])
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
		gr.tiles[t].tileImage = 10
	}
	if !gr.tiles[t].isUncovered && !gr.tiles[t].isMine {
		gr.tiles[t].tileImage = gr.tiles[t].neighborMines
		gr.tiles[t].isUncovered = true
	}
}

func (gr *grid) flag(t int) {
	if gr.tiles[t].tileImage == 9 {
		gr.tiles[t].tileImage = 13
	} else if gr.tiles[t].tileImage == 13 {
		gr.tiles[t].tileImage = 9
	} else {
		return
	}
}
