package player

import (
	"blackjack/deck"
	"blackjack/hand"
	"blackjack/wallet"
)

type RegularPlayer struct {
	score int
	hand hand.Hand
	wallet wallet.Wallet
}

func (regularPlayer *RegularPlayer) GetScore() int{
	return regularPlayer.score
}

func (regularPlayer *RegularPlayer) Win(bet int){
	regularPlayer.score++
	regularPlayer.wallet.WonMoney(bet)
}

func (regularPlayer *RegularPlayer) Loose(bet int){
	regularPlayer.score--
	regularPlayer.wallet.LostMoney(bet)
}

func (regularPlayer *RegularPlayer) DrawCard(deck *deck.Deck){
	card := deck.Draw()
	card.IsVisible = true
	regularPlayer.hand.AddCardToHand(card)
}

func (regularPlayer *RegularPlayer) DiscardAllCards(deck *deck.Deck){
	deck.Discard(regularPlayer.hand.GetHandCards())
	regularPlayer.hand.DiscardAllCards()
}

func (regularPlayer *RegularPlayer) GetHandScore() int{
	return regularPlayer.hand.GetHandCardsSum()
}