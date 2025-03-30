package main

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
