package main

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func debugMouse(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, strconv.FormatBool(ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)))
	fmt.Println(ebiten.CursorPosition())
}
