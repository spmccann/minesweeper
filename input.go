package main

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var locationClicked [2]int

func debugMouse(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, strconv.FormatBool(ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)))
	fmt.Println(ebiten.CursorPosition())
}

func positionClicked() {
	var mousePosX, mousePosY int
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) == true {
		ebiten.CursorPosition()
		mousePosX, mousePosY = ebiten.CursorPosition()
		locationClicked[0] = mousePosX
		locationClicked[1] = mousePosY
	}
}
