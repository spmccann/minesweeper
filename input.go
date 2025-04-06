package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	tileClick                   [2]int
	mouseRightPreviouslyPressed bool
	mouseLeftPreviouslyPressed  bool
	tileWhenPressed             [2]int
	grid                        grid
	mouseButtonLeft             bool
	mouseButtonRight            bool
}

func newInput() input {
	return input{
		tileClick:                  [2]int{-2, -2},
		mouseLeftPreviouslyPressed: false,
		mouseButtonRight:           false,
		tileWhenPressed:            [2]int{-3, -3},
	}
}

func (i *input) tileClicked() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	i.tileClick[0] = (mousePosX - i.grid.offset) / i.grid.tileSize
	i.tileClick[1] = (mousePosY - i.grid.offset) / i.grid.tileSize
}

func (i *input) registerPress() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	outBounds := mousePosX < i.grid.offset || mousePosY < i.grid.offset
	if i.mouseButtonLeft && !i.mouseLeftPreviouslyPressed && !outBounds {
		i.tileWhenPressed = [2]int{(mousePosX - i.grid.offset) / i.grid.tileSize, (mousePosY - i.grid.offset) / i.grid.tileSize}
		i.mouseButtonLeft = false
	}
	if i.mouseButtonRight && !i.mouseRightPreviouslyPressed && !outBounds {
		i.tileWhenPressed = [2]int{(mousePosX - i.grid.offset) / i.grid.tileSize, (mousePosY - i.grid.offset) / i.grid.tileSize}
		i.mouseButtonRight = false
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
	i.tileClick = [2]int{-2, -2}
	i.grid = grid
	leftPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rightPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	if leftPressed {
		i.mouseButtonLeft = true
		i.registerPress()
	}
	if rightPressed {
		i.mouseButtonRight = true
		i.registerPress()
	}
	if !leftPressed && i.mouseLeftPreviouslyPressed && i.comparePosition() {
		i.tileClicked()
	}
	if !rightPressed && i.mouseRightPreviouslyPressed && i.comparePosition() {
		i.tileClicked()
	}
	i.mouseLeftPreviouslyPressed = leftPressed
	i.mouseRightPreviouslyPressed = rightPressed
}

func (i *input) debugCursor() {
	fmt.Println(ebiten.CursorPosition())
}
