package main

import (
	"image"
	"image/color"
	"golang.org/x/image/colornames"
	"log"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten"
)
var (
	fontNumberCard      font.Face
	fontSignCard        font.Face
	fontNrMHeight   int
	fontSMHeight   int
)
func init() {
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	fontNumberCard = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	bN, _, _ := fontNumberCard.GlyphBounds('M')
	fontNrMHeight = (bN.Max.Y - bN.Min.Y).Ceil()

	fontSignCard = truetype.NewFace(tt, &truetype.Options{
		Size:    100,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	b, _, _ := fontSignCard.GlyphBounds('M')
	fontSMHeight = (b.Max.Y - b.Min.Y).Ceil()
}

type Card struct {
	Rect  image.Rectangle
	Number  string
	Sign  string
	Color color.RGBA
}

func (c *Card) Draw(dst *ebiten.Image) {
	cardColorBg := color.RGBA{0xff, 0xff, 0xff, 0xff}
	cardColor := colornames.Black
	if(c.Sign == "♦" || c.Sign == "♥"){
		cardColor = colornames.Red
	}
	drawNinePatches(dst, c.Rect, imageSrcRects[imageTypeButton], cardColorBg)

	bounds, _ := font.BoundString(fontSignCard, c.Sign)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := c.Rect.Min.X + (c.Rect.Dx()-w)/2
	y := c.Rect.Max.Y - (c.Rect.Dy()-fontSMHeight)/2
	text.Draw(dst, c.Sign, fontSignCard, x, y, cardColor)

	// Top left number sign
	x1 := c.Rect.Min.X
	y1 := c.Rect.Min.Y + fontNrMHeight
	y2 := c.Rect.Min.Y + fontNrMHeight + fontNrMHeight
	text.Draw(dst, c.Number, fontNumberCard, x1, y1, cardColor)
	text.Draw(dst, c.Sign, fontNumberCard, x1, y2, cardColor)

	//text.Draw(screen, "♠ ♣ ", fontCard, 50, 50, color.Black)
	//text.Draw(screen, "♦ ♥", fontCard, 50, 100, colornames.Red)
}
