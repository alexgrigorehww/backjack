package ui

import "blackjack/deck"

type UI interface {
	RenderCleanTableWithBettingOptions(setBet func(int) error, walletAmount int)
	RenderDeal(deal func() error)
	RenderHitOrStand(hit func() error, stand func() error)
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
