package player

import (
	"blackjack/deck"
	"blackjack/hand"
)

type Dealer struct {
	score int
	hand  *hand.Hand
	Bet   int
}

type SerializableDealer struct {
	Score int
	Hand  *hand.SerializableHand
}

func (dealer *Dealer) Init() {
	dealer.hand = new(hand.Hand)
}

func (dealer *Dealer) GetScore() int {
	return dealer.score
}

func (dealer *Dealer) Win() int {
	dealer.score++
	return 0
}

func (dealer *Dealer) Loose() int {
	return 0
}

func (dealer *Dealer) DrawCard(deck *deck.Deck) *deck.Card {
	card := deck.Draw()
	dealer.hand.AddCardToHand(card)
	if len(dealer.hand.GetHandCards()) != 2 {
		card.IsVisible = true
	}
	return card
}

func (dealer *Dealer) DiscardAllCards(deck *deck.Deck) {
	deck.Discard(dealer.hand.GetHandCards())
	dealer.hand.DiscardAllCards()
}

func (dealer *Dealer) GetHandScore() int {
	return dealer.hand.GetHandCardsSum()
}

func (dealer *Dealer) GetHandScores() []int {
	return dealer.hand.DisplayValues()
}

func (dealer *Dealer) GetCards() []*deck.Card {
	return dealer.hand.GetHandCards()
}

func (_ *Dealer) GetWalletAmount() int {
	return 0
}

func (dealer *Dealer) IsBusted() bool {
	return dealer.hand.GetHandCardsSum() > 21
}

func (dealer *Dealer) IsBlackjack() bool {
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

func (dealer *Dealer) getHandCards() []*deck.Card {
	return dealer.hand.GetHandCards()
}

func (dealer *Dealer) RevealSecondCard() {
	dealer.hand.GetHandCards()[1].IsVisible = true
}

func (dealer *Dealer) GetSerializable() *SerializableDealer {
	serializableDealer := SerializableDealer{
		Score: dealer.score,
		Hand:  dealer.hand.GetSerializable(),
	}
	return &serializableDealer
}

func (serializableDealer *SerializableDealer) Deserialize() *Dealer {
	dealer := Dealer{
		score: serializableDealer.Score,
		hand:  serializableDealer.Hand.Deserialize(),
	}
	return &dealer
}
