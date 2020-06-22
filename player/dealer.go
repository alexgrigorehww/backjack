package player

import (
	"blackjack/deck"
	"blackjack/hand"
)

type Dealer struct {
	score int
	hand hand.Hand
}

func (dealer *Dealer) GetScore() int{
	return dealer.score
}

func (dealer *Dealer) Win(_ int){
	dealer.score++
}

func (dealer *Dealer) Loose(_ int){
	dealer.score--
}

func (dealer *Dealer) DrawCard(deck *deck.Deck){
	card := deck.Draw()
	dealer.hand.AddCardToHand(card)
	if len(dealer.hand.GetHandCards()) != 2{
		card.IsVisible = true
	}
}

func (dealer *Dealer) DiscardAllCards(deck *deck.Deck){
	deck.Discard(dealer.hand.GetHandCards())
	dealer.hand.DiscardAllCards()
}

func (dealer *Dealer) GetHandScore() int{
	return dealer.hand.GetHandCardsSum()
}

func (dealer *Dealer) GetCards() []*deck.Card{
	return dealer.hand.GetHandCards()
}