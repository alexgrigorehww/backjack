package main

import (
	"blackjack/deck"
	"blackjack/wallet"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
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
	poli    *ebiten.Image
	theDeck *deck.Deck
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
	score         = &Score{20}
	sumCards      = &SumCards{14}
	playerCards   []drawnCard
	dealerCards   []drawnCard
	playerStopped = false
)

type drawnCard struct {
	Card     deck.Card
	IsHidden bool
}

func init() {
	theDeck = new(deck.Deck)
	theDeck.Init()
	theDeck.Shuffle(deck.ShuffleAndMixAll)

	fmt.Sprintf("%+q", 1)
	dealerCards = []drawnCard{
		drawnCard{*theDeck.Draw(), true},
		drawnCard{*theDeck.Draw(), false},
	}

	playerCards = []drawnCard{
		drawnCard{*theDeck.Draw(), false},
		drawnCard{*theDeck.Draw(), false},
	}

}

func update(screen *ebiten.Image) error {
	buttonStand.Update()
	buttonHit.Update()
	buttonHit.SetOnPressed(func(b *Button) {
		if !playerStopped {
			playerCards = append(playerCards, drawnCard{*theDeck.Draw(), false})
		}
	})
	buttonStand.SetOnPressed(func(b *Button) {
		playerStopped = true
	})
	w := new(wallet.Wallet)
	w.SetAmount(500)
	w.LostMoney(20)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{0xeb, 0xeb, 0xeb, 0xff})

	// dealer cards
	renderCards(screen, dealerCards, 125, 25)

	// player cards
	renderCards(screen, playerCards, 100, 165)

	score.Draw(screen, walletMoney)
	buttonStand.Draw(screen)
	buttonHit.Draw(screen)

	return nil
}

func renderCards(screen *ebiten.Image, cards []drawnCard, startX int, startY int) {
	cardsSum := 0
	for uk, card := range cards {
		if card.IsHidden {
			drawNinePatches(screen, image.Rect(startX+15*uk, startY, startX+75+15*uk, startY+125), imageSrcRects[imageTypeButton], colornames.Darkgreen)
		} else {
			(&Card{
				Rect:   image.Rect(startX+15*uk, startY, startX+75+15*uk, startY+125),
				Number: card.Card.GetDisplayingValue(),
				Sign:   card.Card.GetSymbol(),
			}).Draw(screen)
			cardsSum = cardsSum + card.Card.GetBlackjackValue()
		}
	}
	if cardsSum >= 21 {
		playerStopped = true
	}
	sumCards.Draw(screen, cardsSum, startX, startY)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "BlackJack"); err != nil {
		log.Fatal(err)
	}
}
