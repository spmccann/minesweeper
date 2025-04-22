package main

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/spritesheet.png
var tileSpritesFS embed.FS

//go:embed assets/menusprites.png
var menuSpritesFS embed.FS

//go:embed assets/largemenusprites.png
var largeMenuSpritesFS embed.FS

type graphic struct {
	tileSprites      *ebiten.Image
	menuSprites      *ebiten.Image
	largeMenuSprites *ebiten.Image
	tileImages       []*ebiten.Image
	menuImages       []*ebiten.Image
	menuLargeImages  []*ebiten.Image
	sortTileImages   []*ebiten.Image
}

func (gs *graphic) init() {
	var err error

	tileSpritesFile, err := tileSpritesFS.Open("assets/spritesheet.png")
	if err != nil {
		log.Fatal(err)
	}
	defer tileSpritesFile.Close()
	imgTileSprites, _, err := image.Decode(tileSpritesFile)
	if err != nil {
		log.Fatal(err)
	}
	gs.tileSprites = ebiten.NewImageFromImage(imgTileSprites)

	menuSpritesFile, err := menuSpritesFS.Open("assets/menusprites.png")
	if err != nil {
		log.Fatal(err)
	}
	defer menuSpritesFile.Close()
	imgMenuSprites, _, err := image.Decode(menuSpritesFile)
	if err != nil {
		log.Fatal(err)
	}
	gs.menuSprites = ebiten.NewImageFromImage(imgMenuSprites)

	largeMenuSpritesFile, err := largeMenuSpritesFS.Open("assets/largemenusprites.png")
	if err != nil {
		log.Fatal(err)
	}
	defer largeMenuSpritesFile.Close()
	imgLargeMenuSprites, _, err := image.Decode(largeMenuSpritesFile)
	if err != nil {
		log.Fatal(err)
	}
	gs.largeMenuSprites = ebiten.NewImageFromImage(imgLargeMenuSprites)
}

func (gs *graphic) createTileImages() {
	imgWidth := gs.tileSprites.Bounds().Dx()
	tileSize := 32
	padding := 1
	stride := tileSize + padding
	for x := 0; x < imgWidth; x += stride {
		part := gs.tileSprites.SubImage(image.Rect(x, 0, x+tileSize, tileSize))
		newTile := ebiten.NewImageFromImage(part)
		gs.tileImages = append(gs.tileImages, newTile)
	}
	gs.sortImages()
}

func (gs *graphic) createMenuImages() {
	imgWidth := gs.menuSprites.Bounds().Dx()
	tileSize := 32
	padding := 1
	stride := tileSize + padding
	for x := 0; x < imgWidth; x += stride {
		part := gs.menuSprites.SubImage(image.Rect(x, 0, x+tileSize, tileSize))
		newTile := ebiten.NewImageFromImage(part)
		gs.menuImages = append(gs.menuImages, newTile)
	}
}

func (gs *graphic) createLargeMenuImages() {
	imgWidth := gs.largeMenuSprites.Bounds().Dx()
	tileWidth := 64
	tileHeight := 32
	padding := 1
	stride := tileWidth + padding
	for x := 0; x < imgWidth; x += stride {
		part := gs.largeMenuSprites.SubImage(image.Rect(x, 0, x+tileWidth, tileHeight))
		newTile := ebiten.NewImageFromImage(part)
		gs.menuLargeImages = append(gs.menuLargeImages, newTile)
	}
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
