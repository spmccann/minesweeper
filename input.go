package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	locClick               [2]int
	tileClick              [2]int
	mousePreviouslyPressed bool
	tileWhenPressed        [2]int
	grid                   grid
}

func newInput() input {
	return input{
		locClick:               [2]int{-1, -1},
		tileClick:              [2]int{-2, -2},
		mousePreviouslyPressed: false,
		tileWhenPressed:        [2]int{-3, -3},
	}
}

func (i *input) tileClicked() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	i.locClick[0] = mousePosX
	i.locClick[1] = mousePosY
	i.tileClick[0] = (i.locClick[0] - i.grid.offset) / i.grid.tileSize
	i.tileClick[1] = (i.locClick[1] - i.grid.offset) / i.grid.tileSize
}

func (i *input) clickPress() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	mousePresssed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	outBounds := mousePosX < i.grid.offset || mousePosY < i.grid.offset
	if mousePresssed && !i.mousePreviouslyPressed && !outBounds {
		i.tileWhenPressed = [2]int{(mousePosX - i.grid.offset) / i.grid.tileSize, (mousePosY - i.grid.offset) / i.grid.tileSize}
	}
}

func (i *input) comparePosition() bool {
	mousePosX, mousePosY := ebiten.CursorPosition()
	currentPos := [2]int{(mousePosX - i.grid.offset) / i.grid.tileSize, (mousePosY - i.grid.offset) / i.grid.tileSize}
	if i.tileWhenPressed == currentPos {
		return true
	}
	return false
}

func (i *input) clickRelease(grid grid) {
	i.grid = grid
	mousePresssed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	i.clickPress()
	if !mousePresssed && i.mousePreviouslyPressed && i.comparePosition() {
		i.tileClicked()
	}
	i.mousePreviouslyPressed = mousePresssed
}

func (i *input) debugCursor() {
	fmt.Println(ebiten.CursorPosition())
}
