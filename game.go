package main

import (
	"math/rand/v2"
	"time"
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
	tiles     []tile
	offsetX   int
	offsetY   int
	tileSize  int
	gridSize  int
	flags     int
	flagsLeft int
	gameOver  bool
	gameTime  int
	ticker    *time.Ticker
	win       bool
	lost      bool
}

func newGrid() grid {
	return grid{
		tiles:     []tile{},
		offsetX:   64,
		offsetY:   64,
		tileSize:  32,
		gridSize:  320,
		flags:     10,
		flagsLeft: 10,
		gameTime:  0,
		ticker:    time.NewTicker(1 * time.Second),
	}
}

func (gr *grid) populateGrid() {
	t := newTile()
	id := -1
	for x := gr.offsetX; x <= gr.gridSize; x += gr.tileSize {
		for y := gr.offsetY; y <= gr.gridSize; y += gr.tileSize {
			id += 1
			t.updateTile((x-gr.offsetX)/gr.tileSize, (y-gr.offsetY)/gr.tileSize, id)
			gr.tiles = append(gr.tiles, t)
		}
	}
}

func (gr *grid) generateMines(tileException int) {
	mineLocations := randomNumbers(80, gr.flags, tileException)
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

func randomNumbers(maxNum, count, exclusion int) []int {
	selections := make([]int, 0, count)
	numberPool := make([]int, maxNum)
	for i := range numberPool {
		numberPool[i] = i
	}
	for i := 0; i < count; i++ {
		idx := rand.IntN(len(numberPool))
		value := numberPool[idx]
		if value != exclusion {
			selections = append(selections, value)
		}
		numberPool[idx] = numberPool[len(numberPool)-1]
		numberPool = numberPool[:len(numberPool)-1]
	}
	return selections
}

func (gr *grid) checkGrid(in input, s sound) {
	if gr.gameOver {
		gr.ticker.Stop()
		return
	}
	for t := range gr.tiles {
		if gr.tiles[t].x == in.tileClick[0] && gr.tiles[t].y == in.tileClick[1] {
			if in.mouseButtonLeft {
				if s.enabled && s.soundEffects.click != nil {
					s.soundEffects.click.Rewind()
					s.soundEffects.click.Play()
				}
				if !gr.tileClicked() {
					gr.timer()
					gr.generateMines(t)
					gr.neighborNumbers()
				}
				gr.identifyTileClicked(t, s)
			}
			if in.mouseButtonRight {
				if s.enabled && s.soundEffects.flag != nil {
					s.soundEffects.flag.Rewind()
					s.soundEffects.flag.Play()
				}
				gr.flag(t)
				gr.winCheck(s)
			}
		}
	}
}

func (gr *grid) tileClicked() bool {
	for t := range gr.tiles {
		if gr.tiles[t].isUncovered {
			return true
		}
	}
	return false
}

func (gr *grid) identifyTileClicked(t int, s sound) {
	if gr.tiles[t].isMine {
		gr.revealMines(t)
		gr.wrongFlags()
		if s.enabled && s.soundEffects.lose != nil {
			s.soundEffects.lose.Rewind()
			s.soundEffects.lose.Play()
		}
		gr.tiles[t].tileImage = 11
		gr.lost = true
		gr.gameOver = true
	}
	if !gr.tiles[t].isUncovered && !gr.tiles[t].isMine {
		if gr.tiles[t].neighborMines > 0 {
			gr.tiles[t].tileImage = gr.tiles[t].neighborMines
			gr.tiles[t].isUncovered = true
		} else {
			gr.tiles[t].tileImage = 0
			gr.tiles[t].isUncovered = true
			gr.zeroMines(t)

		}
	}
}

func (gr *grid) flag(t int) {
	if gr.tiles[t].tileImage == 9 {
		gr.tiles[t].tileImage = 12
		gr.tiles[t].isFlagged = false
		gr.flagsLeft += 1
	} else if gr.tiles[t].tileImage == 12 {
		gr.tiles[t].tileImage = 13
		gr.tiles[t].isFlagged = false
	} else if gr.tiles[t].tileImage == 13 && gr.flagsLeft != 0 {
		gr.tiles[t].tileImage = 9
		gr.tiles[t].isFlagged = true
		gr.flagsLeft -= 1

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

func (gr *grid) wrongFlags() {
	for t := range gr.tiles {
		if gr.tiles[t].isFlagged && !gr.tiles[t].isMine {
			gr.tiles[t].tileImage = 7
		}
	}
}

func (gr *grid) countFlags() int {
	numFlags := 0
	for t := range gr.tiles {
		if gr.tiles[t].isFlagged {
			numFlags += 1
		}
	}
	return numFlags
}

func (gr *grid) winCheck(s sound) {
	correctFlags := 0
	for t := range gr.tiles {
		if gr.tiles[t].isFlagged && gr.tiles[t].isMine {
			correctFlags += 1
		}
	}

	if correctFlags == gr.flags {
		gr.win = true
		gr.gameOver = true
		if s.enabled && s.soundEffects.win != nil {
			s.soundEffects.win.Rewind()
			s.soundEffects.win.Play()
		}
	}
}

func (gr *grid) zeroMines(tile int) {
	neighborsId := []int{-10, -1, 8, -9, 1, -8, 9, 10}
	neighborCoord := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 1}, {-1, 1}, {1, 0}, {1, 1}}
	for i := range neighborsId {
		nbTile := []int{gr.tiles[tile].x + neighborCoord[i][0], gr.tiles[tile].y + neighborCoord[i][1]}
		xBounds := nbTile[0] >= 0 && nbTile[0] <= 8
		yBounds := nbTile[1] >= 0 && nbTile[1] <= 8
		if xBounds && yBounds {
			nextTile := gr.tiles[tile+neighborsId[i]]
			if nextTile.neighborMines == 0 && !nextTile.isUncovered {
				gr.tiles[nextTile.id].isUncovered = true
				gr.tiles[nextTile.id].tileImage = 0
				gr.zeroMines(nextTile.id)
			} else {
				gr.tiles[nextTile.id].isUncovered = true
				gr.tiles[nextTile.id].tileImage = nextTile.neighborMines
			}
		}
	}
}

func (gr *grid) timer() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-gr.ticker.C:
				gr.gameTime += 1
			}
		}
	}()
}
