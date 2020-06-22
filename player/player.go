package player

import (
	"blackjack/deck"
)

type Player interface {
	GetScore() int
	Win(int) int
	Lose(int) int
	DrawCard(deck.Deck) deck.Card
	DiscardAllCards(deck *deck.Deck)
	GetHandScore() int
	GetCards() []*deck.Card
}
