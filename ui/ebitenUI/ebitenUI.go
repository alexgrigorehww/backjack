package ebitenui

import (
	"blackjack/deck"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

// EbitenUI is the UI using ebiten package
type EbitenUI struct {
	dealerCards []*deck.Card
	playerCards []*deck.Card
	dealerSums  []int
	playerSums  []int

	buttonNewGame *Button
	buttonStand   *Button
	buttonHit     *Button
	buttonSave    *Button
	buttonRestore *Button
	buttonSetBet1 *Button
	buttonSetBet2 *Button
	buttonSetBet3 *Button
	buttonSetBet4 *Button
	buttonDeal    *Button
	score         *Score
	sumCards      *SumCards
	bustText      *RenderText
	cardSum       map[string]int

	walletAmount int
	status       string
}

func (ui *EbitenUI) newGame() {
	ui.dealerCards = nil
	ui.playerCards = nil
	ui.status = ""
}

func (ui *EbitenUI) renderBustText(screen *ebiten.Image) {
	renderBustText := true
	text := ""
	switch ui.status {
	case "dealer_wins":
		text = "Dealer Won"
	case "game_over":
		text = "Game over! Out of money"
	case "draw":
		text = "Draw"
	case "player_busted":
		text = "Player Busted!"
	case "player_won":
		text = "You Won!"
	default:
		renderBustText = false
	}
	if renderBustText {
		startX, startY := 100, 250
		ui.bustText.Draw(screen, text, startX+75/2, startY+125/2)
		ui.buttonNewGame.Draw(screen)
	}

}

func (ui *EbitenUI) cleanUi() {
	ui.dealerCards = nil
	ui.dealerSums = nil
	ui.playerCards = nil
	ui.playerSums = nil
}

func (ui *EbitenUI) RenderPlayerWins() {
	ui.status = "player_won"
}

func (ui *EbitenUI) RenderPlayerBusted() {
	ui.status = "player_busted"
}

func (ui *EbitenUI) RenderDealerWins() {
	ui.status = "dealer_wins"
}

func (ui *EbitenUI) RenderDraw() {
	ui.status = "draw"
}

func (ui *EbitenUI) RenderGameOver() {
	ui.status = "game_over"
}

func (ui *EbitenUI) RenderDealerCards(handSums []int) {
	if handSums != nil {
		ui.dealerSums = handSums
	}
}

func (ui *EbitenUI) RenderPlayerCards() {
	// do nothing because the UI is auto rendering
}

func (ui *EbitenUI) AddPlayerCard(card *deck.Card, playerSums []int) {
	ui.playerCards = append(ui.playerCards, card)
	ui.playerSums = playerSums
}
func (ui *EbitenUI) AddDealerCard(card *deck.Card, dealerSums []int) {
	ui.dealerCards = append(ui.dealerCards, card)
	ui.dealerSums = dealerSums
}

func (ui *EbitenUI) RenderHitOrStand() {
	ui.status = "hit_or_stand"
}

func (ui *EbitenUI) RenderCleanTableWithBettingOptions(walletAmount int) {
	ui.cleanUi()
	ui.walletAmount = walletAmount
	ui.status = "set_bet"
}

func (ui *EbitenUI) RenderDeal() {
	ui.status = "deal"
}

func (ui *EbitenUI) SetGameplayActions(setBet func(int) error, saveGame func(chan error), restoreGame func(chan error), newGame func() error, deal func() error, hit func() error, stand func() error) {
	ui.buttonNewGame.SetOnPressed(func(b *Button) {
		err := newGame()
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonSetBet1.SetOnPressed(func(b *Button) {
		err := setBet(5)
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonSetBet2.SetOnPressed(func(b *Button) {
		err := setBet(10)
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonSetBet3.SetOnPressed(func(b *Button) {
		err := setBet(20)
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonSetBet4.SetOnPressed(func(b *Button) {
		err := setBet(50)
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonSave.SetOnPressed(func(b *Button) {
		ch := make(chan error)
		go saveGame(ch)
		err := <-ch
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonRestore.SetOnPressed(func(b *Button) {
		ch := make(chan error)
		go restoreGame(ch)
		err := <-ch
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonDeal.SetOnPressed(func(b *Button) {
		err := deal()
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonHit.SetOnPressed(func(b *Button) {
		err := hit()
		if err != nil {
			log.Fatal(err)
		}
	})
	ui.buttonStand.SetOnPressed(func(b *Button) {
		err := stand()
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (ui *EbitenUI) renderPlayerCards(screen *ebiten.Image) {
	startX, startY := 100, 165
	if len(ui.playerSums) == 0 {
		return
	}
	cardsSum := ui.playerSums[0]
	for uk, card := range ui.playerCards {
		if !card.IsVisible {
			drawNinePatches(screen, image.Rect(startX+15*uk, startY, startX+75+15*uk, startY+125), imageSrcRects[imageTypeButton], colornames.Darkgreen)
		} else {
			(&Card{
				Rect:   image.Rect(startX+15*uk, startY, startX+75+15*uk, startY+125),
				Number: card.GetDisplayingValue(),
				Sign:   card.GetSymbol(),
			}).Draw(screen)
		}
	}
	ui.sumCards.Draw(screen, cardsSum, startX, startY)
}

func (ui *EbitenUI) renderDealerCards(screen *ebiten.Image) {
	startX, startY := 125, 25
	if len(ui.dealerSums) == 0 {
		return
	}
	cardsSum := ui.dealerSums[0]
	for uk, card := range ui.dealerCards {
		if !card.IsVisible {
			drawNinePatches(screen, image.Rect(startX+15*uk, startY, startX+75+15*uk, startY+125), imageSrcRects[imageTypeButton], colornames.Darkgreen)
		} else {
			(&Card{
				Rect:   image.Rect(startX+15*uk, startY, startX+75+15*uk, startY+125),
				Number: card.GetDisplayingValue(),
				Sign:   card.GetSymbol(),
			}).Draw(screen)
		}
	}

	ui.sumCards.Draw(screen, cardsSum, startX, startY)
}

func (ui *EbitenUI) renderHitOrStand(screen *ebiten.Image) {
	if ui.status != "hit_or_stand" {
		return
	}
	ui.buttonHit.Draw(screen)
	ui.buttonStand.Draw(screen)
}

func (ui *EbitenUI) renderSetBet(screen *ebiten.Image) {
	if ui.status != "set_bet" {
		return
	}
	ui.buttonSetBet1.Draw(screen)
	ui.buttonSetBet2.Draw(screen)
	ui.buttonSetBet3.Draw(screen)
	ui.buttonSetBet4.Draw(screen)
	ui.buttonSave.Draw(screen)
	ui.buttonRestore.Draw(screen)
}

func (ui *EbitenUI) renderDeal(screen *ebiten.Image) {
	if ui.status != "deal" {
		return
	}
	ui.buttonDeal.Draw(screen)

}

func (ui *EbitenUI) update(screen *ebiten.Image) error {
	ui.buttonNewGame.Update()
	ui.buttonStand.Update()
	ui.buttonHit.Update()
	ui.buttonSetBet1.Update()
	ui.buttonSetBet2.Update()
	ui.buttonSetBet3.Update()
	ui.buttonSetBet4.Update()
	ui.buttonDeal.Update()
	ui.buttonSave.Update()
	ui.buttonRestore.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	//// Fill background
	screen.Fill(color.RGBA{0xeb, 0xeb, 0xeb, 0xff})

	ui.score.Draw(screen, ui.walletAmount)
	ui.renderDealerCards(screen)
	ui.renderPlayerCards(screen)

	ui.renderHitOrStand(screen)
	ui.renderBustText(screen)
	ui.renderSetBet(screen)
	ui.renderDeal(screen)

	return nil
}

// Init initializes the UI
func (ui *EbitenUI) Init() {
	ui.buttonNewGame = &Button{
		Rect:  image.Rect(20, 355, 90, 385),
		Text:  "New Game",
		Color: color.RGBA{0x88, 0x88, 0x88, 0xff},
	}
	ui.buttonStand = &Button{
		Rect:  image.Rect(110, 350, 195, 390),
		Text:  "STAND",
		Color: color.RGBA{0x88, 0x88, 0x88, 0xff},
	}
	ui.buttonHit = &Button{
		Rect:  image.Rect(205, 350, 290, 390),
		Text:  "HIT",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	ui.buttonSetBet1 = &Button{
		Rect:  image.Rect(50, 250, 100, 285),
		Text:  "5",
		Color: color.RGBA{0x88, 0x88, 0x88, 0xff},
	}
	ui.buttonSetBet2 = &Button{
		Rect:  image.Rect(110, 250, 160, 285),
		Text:  "10",
		Color: color.RGBA{0x88, 0x88, 0x88, 0xff},
	}
	ui.buttonSetBet3 = &Button{
		Rect:  image.Rect(170, 250, 220, 285),
		Text:  "20",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	ui.buttonSetBet4 = &Button{
		Rect:  image.Rect(230, 250, 280, 285),
		Text:  "50",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	ui.buttonDeal = &Button{
		Rect:  image.Rect(230, 150, 330, 185),
		Text:  "Deal!",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	ui.buttonSave = &Button{
		Rect:  image.Rect(320, 100, 390, 130),
		Text:  "Save Game",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	ui.buttonRestore = &Button{
		Rect:  image.Rect(320, 140, 390, 170),
		Text:  "Restore",
		Color: color.RGBA{0xAA, 0xAA, 0xAA, 0xff},
	}
	ui.score = &Score{20}
	ui.sumCards = &SumCards{14}
	ui.bustText = &RenderText{14, colornames.Black}
	ui.cardSum = map[string]int{"dealer": 0, "player": 0}

	if err := ebiten.Run(ui.update, screenWidth, screenHeight, 1, "BlackJack"); err != nil {
		log.Fatal(err)
	}
}
