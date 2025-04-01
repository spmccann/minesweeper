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

type grid struct {
	tiles    [][]tile
	position int
}

func calcTileClicked() {
	xLoc := (locationClicked[0] - 16) / 16
	yLoc := (locationClicked[1] - 16) / 16
	fmt.Println(xLoc, yLoc)
}

func updateTileImage() {

}
