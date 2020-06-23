package player

import (
	"blackjack/deck"
)

type Player interface {
	GetScore() int
	Win(int)
	Loose(int)
	DrawCard(*deck.Deck)
	DiscardAllCards(*deck.Deck)
	GetHandScore() int
	GetCards() []*deck.Card
	GetWalletAmount() int
}
