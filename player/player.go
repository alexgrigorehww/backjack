package player

import (
	"blackjack/deck"
)

type Player interface {
	GetScore() int
	Win() int
	Loose() int
	DrawCard(*deck.Deck) *deck.Card
	DiscardAllCards(*deck.Deck)
	GetHandScores() []int
	GetHandScore() int
	GetCards() []*deck.Card
	GetWalletAmount() int
	IsBusted() bool
	IsBlackjack() bool
	getHandCards() []*deck.Card
}
