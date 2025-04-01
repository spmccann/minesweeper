package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	current := 16
	maxSize := 144
	for x := current; x <= maxSize; x += current {
		for y := current; y <= maxSize; y += current {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(tileImages[0], op)
		}
	}
	positionClicked()
	calcTileClicked()
	locX, locY := updateTileImage()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("tile %i, %i ", locX, locY))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 176, 176
}

func main() {
	createTileImages(sprites)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Minesweeper")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
