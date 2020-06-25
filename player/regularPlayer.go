package player

import (
	"blackjack/deck"
	"blackjack/hand"
	"blackjack/wallet"
)

type RegularPlayer struct {
	score int
	hand *hand.Hand
	wallet *wallet.Wallet
	Bet int
}

func (regularPlayer *RegularPlayer) Init(){
	regularPlayer.hand = new(hand.Hand)
	regularPlayer.wallet = new(wallet.Wallet)
	regularPlayer.wallet.SetAmount(100)
}

func (regularPlayer *RegularPlayer) GetScore() int{
	return regularPlayer.score
}

func (regularPlayer *RegularPlayer) Win() int{
	regularPlayer.score++
	return regularPlayer.wallet.WonMoney(regularPlayer.Bet)
}

func (regularPlayer *RegularPlayer) Loose() int{
	regularPlayer.score--
	return regularPlayer.wallet.LostMoney(regularPlayer.Bet)
}

func (regularPlayer *RegularPlayer) DrawCard(deck *deck.Deck) *deck.Card{
	card := deck.Draw()
	card.IsVisible = true
	regularPlayer.hand.AddCardToHand(card)
	return card
}

func (regularPlayer *RegularPlayer) DiscardAllCards(deck *deck.Deck){
	deck.Discard(regularPlayer.hand.GetHandCards())
	regularPlayer.hand.DiscardAllCards()
}

func (regularPlayer *RegularPlayer) GetHandScore() int{
	return regularPlayer.hand.GetHandCardsSum()
}

func (regularPlayer *RegularPlayer) GetHandScores() []int{
	return regularPlayer.hand.DisplayValues()
}

func (regularPlayer *RegularPlayer) GetCards() []*deck.Card{
	return regularPlayer.hand.GetHandCards()
}

func (regularPlayer *RegularPlayer) GetWalletAmount() int{
	return regularPlayer.wallet.GetAmount()
}

func (regularPlayer *RegularPlayer)IsBusted() bool{
	busted := true
	handScores := regularPlayer.hand.DisplayValues()

	for _, score := range handScores {
		if score < 21{
			busted = false
			break
		}
	}
	return busted
}

func (regularPlayer *RegularPlayer)IsBlackjack() bool{
	isBlackjack := false
	handScores := regularPlayer.hand.DisplayValues()

	for _, score := range handScores {
		if score == 21 {
			isBlackjack = true
			break
		}
	}
	return isBlackjack
}