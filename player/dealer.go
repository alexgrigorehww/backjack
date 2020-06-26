package player

import (
	"blackjack/deck"
	"blackjack/hand"
)

type Dealer struct {
	score int
	hand *hand.Hand
}

func (dealer *Dealer) Init(){
	dealer.hand = new(hand.Hand)
}

func (dealer *Dealer) GetScore() int{
	return dealer.score
}

func (dealer *Dealer) Win() int{
	dealer.score++
	return 0
}

func (dealer *Dealer) Loose() int{
	dealer.score--
	return 0
}

func (dealer *Dealer) DrawCard(deck *deck.Deck) *deck.Card{
	card := deck.Draw()
	dealer.hand.AddCardToHand(card)
	if len(dealer.hand.GetHandCards()) != 2{
		card.IsVisible = true
	}
	return card
}

func (dealer *Dealer) DiscardAllCards(deck *deck.Deck){
	deck.Discard(dealer.hand.GetHandCards())
	dealer.hand.DiscardAllCards()
}

func (dealer *Dealer) GetHandScore() int{
	return dealer.hand.GetHandCardsSum()
}

func (dealer *Dealer) GetHandScores() []int{
	return dealer.hand.DisplayValues()
}

func (dealer *Dealer) GetCards() []*deck.Card{
	return dealer.hand.GetHandCards()
}

func (_ *Dealer) GetWalletAmount() int{
	return 0
}

func (dealer *Dealer)IsBusted() bool{
	return dealer.hand.GetHandCardsSum() > 21
}

func (dealer *Dealer) IsBlackjack() bool{
	isBlackjack := false
	handScores := dealer.hand.DisplayValues()

	for _, score := range handScores {
		if score == 21 {
			isBlackjack = true
			break
		}
	}
	return isBlackjack
}

func (dealer *Dealer) getHandCards() []*deck.Card{
	return dealer.hand.GetHandCards()
}