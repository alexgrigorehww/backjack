package main

import (
	"fmt"
	"image/color"
	"log"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
)

var (
	arcadeFont font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}
type Score struct {

}

func (s *Score) Draw(dst *ebiten.Image) {
	scoreStr := fmt.Sprintf("%04d", walletMoney)
	text.Draw(dst, scoreStr, arcadeFont, screenWidth-len(scoreStr)*fontSize, fontSize, color.White)
}