package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	tileClick                   [2]int
	menuClick                   [2]int
	mouseRightPreviouslyPressed bool
	mouseLeftPreviouslyPressed  bool
	tileWhenPressed             [2]int
	menuWhenPressed             [2]int
	grid                        grid
	menu                        menu
	mouseButtonLeft             bool
	mouseButtonRight            bool
}

func newInput() input {
	return input{
		tileClick:                  [2]int{-2, -2},
		menuClick:                  [2]int{-1, -1},
		mouseLeftPreviouslyPressed: false,
		mouseButtonRight:           false,
		tileWhenPressed:            [2]int{-3, -3},
	}
}

func (i *input) tileClicked() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	i.tileClick[0] = (mousePosX - i.grid.offsetX) / i.grid.tileSize
	i.tileClick[1] = (mousePosY - i.grid.offsetY) / i.grid.tileSize
}

func (i *input) menuClicked() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	i.menuClick[0] = (mousePosX - i.menu.offsetX) / i.menu.itemWidth
	i.menuClick[1] = (mousePosY - i.menu.offsetY) / i.menu.itemHeight
}
func (i *input) registerPress() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	outBounds := mousePosX < i.grid.offsetX || mousePosY < i.grid.offsetY
	if i.mouseButtonLeft && !i.mouseLeftPreviouslyPressed && !outBounds {
		i.tileWhenPressed = [2]int{(mousePosX - i.grid.offsetX) / i.grid.tileSize, (mousePosY - i.grid.offsetY) / i.grid.tileSize}
		i.mouseButtonLeft = false
	}
	if i.mouseButtonRight && !i.mouseRightPreviouslyPressed && !outBounds {
		i.tileWhenPressed = [2]int{(mousePosX - i.grid.offsetX) / i.grid.tileSize, (mousePosY - i.grid.offsetY) / i.grid.tileSize}
		i.mouseButtonRight = false
	}
	if i.mouseButtonLeft && !i.mouseLeftPreviouslyPressed && outBounds {
		i.menuWhenPressed = [2]int{(mousePosX - i.menu.offsetX) / i.menu.itemWidth, (mousePosY - i.menu.offsetY) / i.menu.itemHeight}
		i.mouseButtonLeft = false
	}
}

func (i *input) comparePosition() bool {
	mousePosX, mousePosY := ebiten.CursorPosition()
	currentPos := [2]int{(mousePosX - i.grid.offsetX) / i.grid.tileSize, (mousePosY - i.grid.offsetY) / i.grid.tileSize}
	if i.tileWhenPressed == currentPos {
		return true
	}
	if i.menuWhenPressed == currentPos {
		return true
	}
	return false
}

func (i *input) mouseEvents(grid grid, menu menu) {
	i.tileClick = [2]int{-2, -2}
	i.menuClick = [2]int{-1, -1}
	i.grid = grid
	i.menu = menu
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
		i.menuClicked()
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
