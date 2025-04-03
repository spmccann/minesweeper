package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	tileClick              [2]int
	mousePreviouslyPressed bool
	tileWhenPressed        [2]int
	grid                   grid
	button                 ebiten.MouseButton
}

func newInput() input {
	return input{
		tileClick:              [2]int{-2, -2},
		mousePreviouslyPressed: false,
		tileWhenPressed:        [2]int{-3, -3},
	}
}

func (i *input) tileClicked() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	i.tileClick[0] = (mousePosX - i.grid.offset) / i.grid.tileSize
	i.tileClick[1] = (mousePosY - i.grid.offset) / i.grid.tileSize
}

func (i *input) clickPress(button ebiten.MouseButton) {
	mousePosX, mousePosY := ebiten.CursorPosition()
	mousePresssed := ebiten.IsMouseButtonPressed(button)
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

func (i *input) clickRelease(grid grid, button ebiten.MouseButton) {
	i.grid = grid
	mousePresssed := ebiten.IsMouseButtonPressed(button)
	i.clickPress(button)
	if !mousePresssed && i.mousePreviouslyPressed && i.comparePosition() {
		i.tileClicked()
		i.button = button
	}
	i.mousePreviouslyPressed = mousePresssed
}

func (i *input) debugCursor() {
	fmt.Println(ebiten.CursorPosition())
}
