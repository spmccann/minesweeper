package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	newGame bool
	grid    grid
	input   input
	graphic graphic
	menu    menu
}

func (g *Game) Update() error {
	if g.newGame {
		g.input = newInput()

		g.grid = newGrid()
		g.grid.populateGrid()

		g.menu = newMenu()
		g.menu.populateMenu()

		g.graphic.init()
		g.graphic.createTileImages()
		g.graphic.createMenuImages()

		g.newGame = false
	}
	g.input.mouseEvents(g.grid, g.menu)
	g.grid.checkGrid(g.input)
	g.menu.checkMenu(g.input)
	if g.menu.items[0].onSelect == true {
		g.newGame = true
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.displayMenu(screen)
	g.displayGrid(screen)
}

func (g *Game) displayMenu(screen *ebiten.Image) {
	i := 0
	for x := g.menu.offsetX; x <= g.menu.menuWidth; x += g.menu.itemWidth {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(0))
		imgNum := g.menu.items[i].itemImage
		screen.DrawImage(g.graphic.sortTileImages[imgNum], op)
		i += 1
	}
}

func (g *Game) displayGrid(screen *ebiten.Image) {
	i := 0
	for x := g.grid.offsetX; x <= g.grid.gridSize; x += g.grid.tileSize {
		for y := g.grid.offsetY; y <= g.grid.gridSize; y += g.grid.tileSize {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			imgNum := g.grid.tiles[i].tileImage
			screen.DrawImage(g.graphic.sortTileImages[imgNum], op)
			i += 1
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 480
}

func windowIcon() {
	f, err := os.Open("assets/new_assets/Flag.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowIcon([]image.Image{img})
}

func main() {
	windowIcon()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Minesweeper")
	if err := ebiten.RunGame(&Game{newGame: true}); err != nil {
		log.Fatal(err)
	}
}
