package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

var sprites *ebiten.Image
var tileImages []*ebiten.Image

func init() {
	var err error
	sprites, _, err = ebitenutil.NewImageFromFile("assets/sprite_sheet.png")
	if err != nil {
		log.Fatal(err)
	}
}

func createTileImages(img *ebiten.Image) {
	for x := 0; x <= 48; x += 16 {
		for y := 0; y <= 48; y += 16 {
			part := img.SubImage(image.Rect(x, y, x+16, y+16))
			newTile := ebiten.NewImageFromImage(part)
			tileImages = append(tileImages, newTile)
		}
	}
}
