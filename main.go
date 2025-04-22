package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"os"
)

type Game struct {
	newGame bool
	grid    grid
	input   input
	graphic graphic
	menu    menu
	sound   sound
}

func (g *Game) Update() error {
	if g.newGame {
		g.input = newInput()

		g.grid = newGrid()
		g.grid.populateGrid()

		g.menu = newMenu()
		g.menu.populateMenu()

		g.graphic.init()
		g.graphic.createMenuImages()
		g.graphic.createLargeMenuImages()
		g.graphic.createTileImages()

		g.sound.init()

		g.newGame = false
	}
	g.input.mouseEvents(g.grid, g.menu)
	g.grid.checkGrid(g.input, g.sound)
	g.menu.checkMenu(g.input)
	g.menu.flagCounter(g.grid.flagsLeft)
	g.menu.timerDisplay(g.grid.gameTime)
	if g.grid.win {
		g.menu.populateLargeMenu(0)
	}
	if g.grid.lost {
		g.menu.populateLargeMenu(1)
	}
	if g.menu.items[0].onSelect == true {
		g.newGame = true
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.displayMenu(screen)
	g.displayLargeMenu(screen)
	g.displayGrid(screen)
}

func (g *Game) displayMenu(screen *ebiten.Image) {
	x := 64
	for t := range g.menu.items {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(0))
		imgNum := g.menu.items[t].itemImage
		screen.DrawImage(g.graphic.menuImages[imgNum], op)
		x += 32
	}
}

func (g *Game) displayLargeMenu(screen *ebiten.Image) {
	x := 320
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(0))
	var imgNum int
	if len(g.menu.largeItems) > 0 {
		imgNum = g.menu.largeItems[0].itemImage
		screen.DrawImage(g.graphic.menuLargeImages[imgNum], op)
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
	if err := ebiten.RunGame(&Game{newGame: true, sound: newSound()}); err != nil {
		log.Fatal(err)
	}
}
