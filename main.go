package main

import (
	"log"
	_ "image/jpeg"
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"git/osero/osero"

)

const (
	screenWidth  = 500	//幅
	screenHeight = 355
)

var (
	Test *ebiten.Image
	mplusNormalFont font.Face
)

func init() {
	var err error
	Test, _, err = ebitenutil.NewImageFromFile("test.jpeg", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    39,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func main(){
	g := game.Game{}
	update:=g.update()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Osero start"); err != nil {
		log.Fatal(err)
	}
}