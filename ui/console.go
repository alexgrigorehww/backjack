package ui

import (
	"blackjack/deck"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ConsoleUi struct {
	dealerCards []*deck.Card
	dealerSums  []int

	playerCards []*deck.Card
	playerSums  []int

	setBet      func(int) error
	saveGame    func(chan error)
	restoreGame func(chan error)
	newGame     func() error
	deal        func() error
	hit         func() error
	stand       func() error
}

func (consoleUi *ConsoleUi) cleanUi() {
	consoleUi.dealerCards = nil
	consoleUi.dealerSums = nil
	consoleUi.playerCards = nil
	consoleUi.playerSums = nil
}

func (consoleUi *ConsoleUi) SetGameplayActions(setBet func(int) error, saveGame func(chan error), restoreGame func(chan error), newGame func() error, deal func() error, hit func() error, stand func() error) {
	consoleUi.setBet = setBet
	consoleUi.saveGame = saveGame
	consoleUi.restoreGame = restoreGame
	consoleUi.newGame = newGame
	consoleUi.deal = deal
	consoleUi.hit = hit
	consoleUi.stand = stand
}

func (consoleUi *ConsoleUi) RenderCleanTableWithBettingOptions(walletAmount int) {
	fmt.Println("New Game! Your wallet :" + strconv.Itoa(walletAmount))
	option := read("You can save (s) / restore (r) or choose your bet")
	var err error
	ch := make(chan error)
	switch option {
	case "s":
		go consoleUi.saveGame(ch)
		err = <-ch
		if err != nil {
			fmt.Println(err)
		}
	case "r":
		go consoleUi.restoreGame(ch)
		err = <-ch
		if err != nil {
			fmt.Println(err)
		}
	default:
		bet, err := strconv.Atoi(option)
		if err != nil {
			err = errors.New("invalid option")
			consoleUi.RenderCleanTableWithBettingOptions(walletAmount)
		} else {
			consoleUi.cleanUi()
			err = consoleUi.setBet(bet)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (consoleUi *ConsoleUi) RenderDeal() {
	dealRes := read("You want to deal? (y/n)")
	if dealRes == "y" {
		consoleUi.deal()
	} else {
		consoleUi.RenderDeal()
	}
}

func (consoleUi *ConsoleUi) RenderHitOrStand() {
	consoleUi.RenderDealerCards(nil)
	consoleUi.RenderPlayerCards()
	action := read("HIT (h) / STAND (s)")
	if action == "h" {
		err := consoleUi.hit()
		if err != nil {
			fmt.Println(err)
		}
	} else if action == "s" {
		err := consoleUi.stand()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		consoleUi.RenderHitOrStand()
	}
}

func (consoleUi *ConsoleUi) AddPlayerCard(card *deck.Card, playerSums []int) {
	consoleUi.playerCards = append(consoleUi.playerCards, card)
	consoleUi.playerSums = playerSums
}

func (consoleUi *ConsoleUi) AddDealerCard(card *deck.Card, dealerSums []int) {
	consoleUi.dealerCards = append(consoleUi.dealerCards, card)
	consoleUi.dealerSums = dealerSums
}

func (consoleUi *ConsoleUi) RenderPlayerBusted() {
	println("Player busted")
	consoleUi.newGame()
}

func (consoleUi *ConsoleUi) RenderPlayerWins() {
	println("Player wins!")
	consoleUi.newGame()
}

func (consoleUi *ConsoleUi) RenderDraw() {
	println("DRAW")
	consoleUi.newGame()
}

func (consoleUi *ConsoleUi) RenderDealerWins() {
	println("Dealer wins")
	consoleUi.newGame()
}

func (consoleUi *ConsoleUi) RenderDealerCards(handSum []int) {
	displayingValues := ""
	for _, card := range consoleUi.dealerCards {
		if !card.IsVisible {
			displayingValues += " ? "
		} else {
			displayingValues += " " + card.GetDisplayingValue() + " "
		}
	}
	if handSum == nil {
		handSum = consoleUi.dealerSums
	}
	fmt.Printf("Dealer cards: %s Sum: %d \n", displayingValues, handSum)
}

func (consoleUi *ConsoleUi) RenderPlayerCards() {
	displayingValues := ""
	for _, card := range consoleUi.playerCards {
		if !card.IsVisible {
			displayingValues += " ? "
		} else {
			displayingValues += " " + card.GetDisplayingValue() + " "
		}
	}
	fmt.Printf("Player cards: %s Sum: %d \n", displayingValues, consoleUi.playerSums)
}

func (consoleUi *ConsoleUi) RenderGameOver() {
	fmt.Print("Game Over! Your wallet is empty")
}

func read(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label + ": ")
	text, _ := reader.ReadString('\n')
	return strings.Trim(text, "\n\r")
}
