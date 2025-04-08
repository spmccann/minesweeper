package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	newGame bool
	grid    grid
	input   input
	graphic graphic
}

func (g *Game) Update() error {
	if g.newGame {
		g.input = newInput()

		g.grid = newGrid()

		g.grid.populateGrid()

		g.graphic.init()
		g.graphic.createTileImages()

		g.newGame = false
	}
	g.input.mouseEvents(g.grid)
	g.grid.checkGrid(g.input)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.displayMenu(screen)
	g.displayGrid(screen)
}

func (g *Game) displayMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0), float64(0))
	screen.DrawImage(g.graphic.sortTileImages[1], op)
}

func (g *Game) displayGrid(screen *ebiten.Image) {
	i := 0
	for x := g.grid.offset; x <= g.grid.gridSize; x += g.grid.tileSize {
		for y := g.grid.offset; y <= g.grid.gridSize; y += g.grid.tileSize {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			imgNum := g.grid.tiles[i].tileImage
			screen.DrawImage(g.graphic.sortTileImages[imgNum], op)
			i += 1
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 416, 416
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Minesweeper")
	if err := ebiten.RunGame(&Game{newGame: true}); err != nil {
		log.Fatal(err)
	}
}
