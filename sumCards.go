package main

import (
	"fmt"
	"image/color"
	"log"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
)

var (
	sumCardsFont font.Face
)

func init() {
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	sumCardsFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}
type SumCards struct {
	fontSize int
}

func (s *SumCards) Draw(dst *ebiten.Image, sumCards int, startX int, startY int) {
	nrStr := fmt.Sprintf("%2d", sumCards)
	text.Draw(dst, nrStr, sumCardsFont, startX - len(nrStr)*s.fontSize, startY + s.fontSize, color.Black)
}