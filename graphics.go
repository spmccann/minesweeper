package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type graphic struct {
	sprites    *ebiten.Image
	tileImages []*ebiten.Image
}

func (gs *graphic) init() {
	var err error
	gs.sprites, _, err = ebitenutil.NewImageFromFile("assets/spritepad.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *graphic) createTileImages() {
	imgWidth := gs.sprites.Bounds().Dx()
	tileSize := 16
	padding := 0
	stride := tileSize + padding
	fmt.Println(imgWidth)
	for x := 0; x < imgWidth; x += 16 {
		part := gs.sprites.SubImage(image.Rect(x, 0, x+stride, 16))
		newTile := ebiten.NewImageFromImage(part)
		gs.tileImages = append(gs.tileImages, newTile)
		padding = 1
	}
}
