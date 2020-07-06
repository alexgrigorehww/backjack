package ui

import "blackjack/deck"

type UI interface {
	RenderCleanTableWithBettingOptions(walletAmount int)
	RenderDeal()
	RenderHitOrStand()
	SetGameplayActions(setBet func(int) error, saveGame func(chan error), restoreGame func(chan error), newGame func() error, deal func() error, hit func() error, stand func() error)
	AddPlayerCard(card *deck.Card, playerSums []int)
	AddDealerCard(card *deck.Card, dealerSums []int)
	RenderPlayerWins()
	RenderPlayerBusted()
	RenderDealerWins()
	RenderDraw()
	RenderGameOver()
	RenderDealerCards(handSum []int)
	RenderPlayerCards()
}
