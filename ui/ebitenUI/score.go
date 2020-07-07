package ebitenui

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var (
	arcadeFont font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

type Score struct {
	fontSize int
}

func (s *Score) Draw(dst *ebiten.Image, walletMoney int) {
	scoreStr := fmt.Sprintf("%04d", walletMoney)
	text.Draw(dst, scoreStr, arcadeFont, screenWidth-len(scoreStr)*s.fontSize, s.fontSize+30, color.Black)
}
