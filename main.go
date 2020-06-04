package main

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/color"
	"log"
)

const (
	screenWidth  = 400
	screenHeight = 400
	fontSize     = 32
	walletMoney  = 500
)

var (
	buttonStand = &Button{
		Rect: image.Rect(110, 350, 195, 390),
		Text: "STAND",
	}

	buttonHit = &Button{
		Rect: image.Rect(205, 350, 290, 390),
		Text: "HIT",
	}
)

func update(screen *ebiten.Image) error {
	buttonStand.Update()
	buttonHit.Update()
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{0xeb, 0xeb, 0xeb, 0xff})

	new(Score).Draw(screen)

	buttonStand.Draw(screen)
	buttonHit.Draw(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "BlackJack"); err != nil {
		log.Fatal(err)
	}
}
