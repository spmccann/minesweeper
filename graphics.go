package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

type graphic struct {
	sprites    *ebiten.Image
	tileImages []*ebiten.Image
}

func (gs *graphic) init() {
	var err error
	gs.sprites, _, err = ebitenutil.NewImageFromFile("assets/sprite_sheet.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *graphic) createTileImages() {
	for x := 0; x <= 48; x += 16 {
		for y := 0; y <= 48; y += 16 {
			part := gs.sprites.SubImage(image.Rect(x, y, x+16, y+16))
			newTile := ebiten.NewImageFromImage(part)
			gs.tileImages = append(gs.tileImages, newTile)
		}
	}
}
