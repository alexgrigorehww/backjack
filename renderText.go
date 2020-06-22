package main

import (
	"image/color"
	"log"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
)

var (
	renderTextFont font.Face
)

func init() {
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	renderTextFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

type RenderText struct {
	fontSize int
	color color.RGBA
}

func (s *RenderText) Draw(dst *ebiten.Image, renderText string, startX int, startY int) {
	text.Draw(dst, renderText, renderTextFont, startX, startY, s.color)
}
