package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	newGame := true
	var tileData grid
	if newGame {
		tileData = updateTileData(newGame, tileData)
		newGame = false
	} else {
		tileData = updateTileData(newGame, tileData)
	}
	positionClicked()
	calcTileClicked()
	current := 16
	maxSize := 144
	i := 0
	for x := current; x <= maxSize; x += current {
		for y := current; y <= maxSize; y += current {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(tileImages[tileData.tiles[i].tileImage], op)
			i += 1
		}
	}
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
