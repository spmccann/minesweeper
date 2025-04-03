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

func (i *input) registerPress(button ebiten.MouseButton) {
	mousePosX, mousePosY := ebiten.CursorPosition()
	mousePressed := ebiten.IsMouseButtonPressed(button)
	outBounds := mousePosX < i.grid.offset || mousePosY < i.grid.offset
	if mousePressed && !i.mousePreviouslyPressed && !outBounds {
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

func (i *input) mouseEvents(grid grid) {
	var mousePressed bool
	i.grid = grid
	leftPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rightPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	if leftPressed {
		mousePressed = leftPressed
		i.button = ebiten.MouseButtonLeft
	}
	if rightPressed {
		mousePressed = rightPressed
		i.button = ebiten.MouseButtonRight
	}
	i.registerPress(i.button)
	if !mousePressed && i.mousePreviouslyPressed && i.comparePosition() {
		i.tileClicked()
	}
	i.mousePreviouslyPressed = mousePressed
}

func (i *input) debugCursor() {
	fmt.Println(ebiten.CursorPosition())
}
