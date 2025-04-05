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
	mouseButtonLeft        bool
	mouseButtonRight       bool
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

func (i *input) registerPress(button bool) {
	var mousePressed bool
	mousePosX, mousePosY := ebiten.CursorPosition()
	if button {
		mousePressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	} else {
		mousePressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	}
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
		i.mouseButtonLeft = true
	}
	if rightPressed {
		mousePressed = rightPressed
		i.mouseButtonRight = true
	}
	i.registerPress(i.mouseButtonLeft)
	i.registerPress(i.mouseButtonRight)
	if !mousePressed && i.mousePreviouslyPressed && i.comparePosition() {
		i.tileClicked()
	}
	i.mousePreviouslyPressed = mousePressed
}

func (i *input) debugCursor() {
	fmt.Println(ebiten.CursorPosition())
}
