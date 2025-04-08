package main

import (
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type graphic struct {
	sprites        *ebiten.Image
	tileImages     []*ebiten.Image
	sortTileImages []*ebiten.Image
}

func (gs *graphic) init() {
	var err error
	gs.sprites, _, err = ebitenutil.NewImageFromFile("assets/spritesheet.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *graphic) createTileImages() {
	imgWidth := gs.sprites.Bounds().Dx()
	tileSize := 32
	padding := 1
	stride := tileSize + padding
	for x := 0; x < imgWidth; x += stride {
		part := gs.sprites.SubImage(image.Rect(x, 0, x+tileSize, tileSize))
		newTile := ebiten.NewImageFromImage(part)
		gs.tileImages = append(gs.tileImages, newTile)
	}
	gs.sortImages()
}

func (gs *graphic) sortImages() {
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[9])  // alt blank
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[0])  // 1
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[1])  // 2
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[2])  // 3
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[8])  // 4
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[10]) // 5
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[7])  // 6
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[6])  // 7
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[3])  // 8
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[5])  // flag
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[11]) //mine
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[13]) //exploded mine
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[12]) // question
	gs.sortTileImages = append(gs.sortTileImages, gs.tileImages[4])  // blank
}
