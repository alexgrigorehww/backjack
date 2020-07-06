package ebitenui

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
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
	color    color.RGBA
}

func (s *RenderText) Draw(dst *ebiten.Image, renderText string, startX int, startY int) {
	text.Draw(dst, renderText, renderTextFont, startX, startY, s.color)
}
