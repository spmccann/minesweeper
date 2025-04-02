package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	locClick               [2]int
	tileClick              [2]int
	mousePreviouslyPressed bool
	tileWhenPressed        [2]int
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
	i.tileClick[0] = (i.locClick[0] - 16) / 16
	i.tileClick[1] = (i.locClick[1] - 16) / 16
}

func (i *input) clickPress() {
	mousePosX, mousePosY := ebiten.CursorPosition()
	mousePresssed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mousePresssed == true && i.mousePreviouslyPressed == false {
		i.tileWhenPressed = [2]int{(mousePosX - 16) / 16, (mousePosY - 16) / 16}
	}
}

func (i *input) comparePosition() bool {
	mousePosX, mousePosY := ebiten.CursorPosition()
	currentPos := [2]int{(mousePosX - 16) / 16, (mousePosY - 16) / 16}
	if i.tileWhenPressed == currentPos {
		return true
	}
	return false
}

func (i *input) clickRelease() {
	mousePresssed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	i.clickPress()
	if !mousePresssed && i.mousePreviouslyPressed && i.comparePosition() {
		i.tileClicked()
	}
	i.mousePreviouslyPressed = mousePresssed
}
