package ui

import (
	"blackjack/deck"
	"bufio"
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
}

func (consoleUi *ConsoleUi) cleanUi() {
	consoleUi.dealerCards = nil
	consoleUi.dealerSums = nil
	consoleUi.playerCards = nil
	consoleUi.playerSums = nil
}

func (consoleUi *ConsoleUi) RenderCleanTableWithBettingOptions(setBet func(int) error, walletAmount int) {
	bet, err := strconv.Atoi(read("New Game! Your wallet: " + strconv.Itoa(walletAmount) + ". Type your bet"))
	if err != nil {
		consoleUi.RenderCleanTableWithBettingOptions(setBet, walletAmount)
	} else {
		consoleUi.cleanUi()
		err := setBet(bet)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (consoleUi *ConsoleUi) RenderDeal(deal func() error) {
	dealRes := read("You want to deal? (y/n)")
	if dealRes == "y" {
		deal()
	} else {
		consoleUi.RenderDeal(deal)
	}
}

func (consoleUi *ConsoleUi) RenderHitOrStand(hit func() error, stand func() error) {
	consoleUi.RenderDealerCards(nil)
	consoleUi.RenderPlayerCards()
	action := read("HIT (h) / STAND (s)")
	if action == "h" {
		err := hit()
		if err != nil {
			fmt.Println(err)
		}
	} else if action == "s" {
		err := stand()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		consoleUi.RenderHitOrStand(hit, stand)
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
}

func (consoleUi *ConsoleUi) RenderPlayerWins() {
	println("Player wins!")
}

func (consoleUi *ConsoleUi) RenderDraw() {
	println("DRAW")
}

func (consoleUi *ConsoleUi) RenderDealerWins() {
	println("Dealer wins")
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
