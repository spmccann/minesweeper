package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	newGame   bool
	tileData  grid
	mouseData input
	graphics  graphic
}

func (g *Game) Update() error {
	if g.newGame {
		g.mouseData = newInput()

		g.tileData = newGrid()

		g.tileData.populateGrid()

		g.graphics.init()
		g.graphics.createTileImages()

		g.newGame = false
	}
	g.mouseData.clickRelease()
	g.tileData.checkGrid(g.mouseData.tileClick)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	current := 16
	maxSize := 144
	i := 0
	for x := current; x <= maxSize; x += current {
		for y := current; y <= maxSize; y += current {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			imgNum := g.tileData.tiles[i].tileImage

			screen.DrawImage(g.graphics.tileImages[imgNum], op)
			i += 1
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 176, 176
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Minesweeper")
	if err := ebiten.RunGame(&Game{newGame: true}); err != nil {
		log.Fatal(err)
	}
}
