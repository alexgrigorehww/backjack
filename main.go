package main

import (
	"blackjack/wallet"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"log"
	"math/rand"
)

const (
	screenWidth  = 400
	screenHeight = 400
	fontSize     = 32
	walletMoney  = 500
)

var (
	poli *ebiten.Image
)
var (
	buttonStand = &Button{
		Rect:  image.Rect(110, 350, 195, 390),
		Text:  "STAND",
		Color: color.RGBA{0x88, 0x88, 0x88, 0xff},
	}
	buttonHit = &Button{
		Rect:  image.Rect(205, 350, 290, 390),
		Text:  "HIT",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	score = &Score{
		fontSize: 20,
	}
	playerCards []drawnCard
	dealerCards []drawnCard
)

type drawnCard struct {
	Number   string
	Sign     string
	IsHidden bool
}

func init() {
	playerCards = []drawnCard{
		drawnCard{"7", "♠", false},
		drawnCard{"K", "♣", false},
		drawnCard{"Q", "♦", false},
		drawnCard{"Q", "♥", false},
	}
	dealerCards = []drawnCard{
		drawnCard{"7", "♠", true},
		drawnCard{"A", "♠", false},
	}
}

func update(screen *ebiten.Image) error {
	buttonStand.Update()
	buttonHit.Update()
	buttonHit.SetOnPressed(func(b *Button) {
		playerCards = append(playerCards, drawnCard{fmt.Sprintf("%d", rand.Intn(9)+2), "♠", false})
	})
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	screen.Fill(color.RGBA{0xeb, 0xeb, 0xeb, 0xff})

	// dealer cards
	renderCards(screen, dealerCards, 25)

	// player cards
	renderCards(screen, playerCards, 165)

	score.Draw(screen)
	buttonStand.Draw(screen)
	buttonHit.Draw(screen)

	return nil
}

func renderCards(screen *ebiten.Image, cards []drawnCard, startY int) {
	for uk, card := range cards {
		if card.IsHidden {
			drawNinePatches(screen, image.Rect(100+15*uk, startY, 175+15*uk, startY+125), imageSrcRects[imageTypeButton], colornames.Darkgreen)
		} else {
			(&Card{
				Rect:   image.Rect(100+15*uk, startY, 175+15*uk, startY+125),
				Number: card.Number,
				Sign:   card.Sign,
			}).Draw(screen)
		}

	}
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "BlackJack"); err != nil {
		log.Fatal(err)
	}
	w := wallet.Wallet{500}
	fmt.Print(w.LostMoney(20))
}
