package gameplay

import (
	"blackjack/deck"
	"blackjack/player"
	"blackjack/ui"
	"errors"
)

const (
	nextStepInit       = ""
	nextStepSetBet     = "SET_BET"
	nextStepDeal       = "DEAL"
	nextStepHitOrStand = "HIT_OR_STAND"
	nextStepNewGame    = "NEW_GAME"
)

type SinglePlayer struct {
	whatsNext string
	dealer    *player.Dealer
	player    *player.RegularPlayer
	deck      *deck.Deck
	consoleUi *ui.ConsoleUi
}

func (gameplay *SinglePlayer) Init() (err error) {
	if gameplay.whatsNext != nextStepInit {
		err = errors.New("invalid gameplay state. cannot initialize the game")
		return
	}
	consoleUi := new(ui.ConsoleUi)

	theDeck := new(deck.Deck)
	theDeck.Init()
	theDeck.Shuffle(deck.ShuffleAndMixAll)

	dealer := new(player.Dealer)
	dealer.Init()

	player := new(player.RegularPlayer)
	player.Init()

	// init UI
	gameplay.consoleUi = consoleUi

	// init gameplay
	gameplay.deck = theDeck
	gameplay.dealer = dealer
	gameplay.player = player

	gameplay.whatsNext = nextStepNewGame
	gameplay.NewGame()

	return
}

func (gameplay *SinglePlayer) SetBet(bet int) (err error) {
	if gameplay.whatsNext != nextStepSetBet {
		err = errors.New("invalid gameplay state. you cannot set bet")
		return
	}
	if gameplay.player.GetWalletAmount() < bet {
		err = errors.New("insufficient amount")
		return
	}
	gameplay.player.Bet = bet
	gameplay.consoleUi.RenderDeal(gameplay.Deal)
	gameplay.whatsNext = nextStepDeal
	return
}

func (gameplay *SinglePlayer) Deal() (err error) {
	if gameplay.whatsNext != nextStepSetBet {
		err = errors.New("invalid gameplay state. you cannot deal")
		return
	}
	playerDrawCard(gameplay)
	dealerDrawCard(gameplay)
	playerDrawCard(gameplay)
	dealerDrawCard(gameplay)

	gameplay.whatsNext = nextStepHitOrStand
	gameplay.consoleUi.RenderHitOrStand(gameplay.Hit, gameplay.Stand)
	return
}

func (gameplay *SinglePlayer) Hit() (err error) {
	if gameplay.whatsNext != nextStepHitOrStand {
		err = errors.New("invalid gameplay state. you cannot hit")
		return
	}
	shouldStop := playerDrawCard(gameplay)
	if !shouldStop {
		gameplay.consoleUi.RenderHitOrStand(gameplay.Hit, gameplay.Stand)
	}
	return
}

func (gameplay *SinglePlayer) Stand() (err error) {
	if gameplay.whatsNext != nextStepHitOrStand {
		err = errors.New("invalid gameplay state. you cannot stand")
		return
	}
	for {
		shouldStop := dealerDrawCard(gameplay)
		if shouldStop {
			break
		}
		dealerScore := gameplay.dealer.GetHandScore()
		if dealerScore >= 17 {
			gameplay.evaluate()
			gameplay.whatsNext = nextStepNewGame
			gameplay.NewGame()
			break
		}
	}
	return
}

func (gameplay *SinglePlayer) NewGame() (err error) {
	if gameplay.whatsNext != nextStepNewGame {
		err = errors.New("invalid gameplay state. you cannot start new game")
		return
	}
	if gameplay.player.GetWalletAmount() == 0 {
		gameplay.consoleUi.RenderGameOver()
	}
	gameplay.dealer.DiscardAllCards(gameplay.deck)
	gameplay.player.DiscardAllCards(gameplay.deck)
	gameplay.whatsNext = nextStepSetBet
	walletAmount := gameplay.player.GetWalletAmount()
	gameplay.consoleUi.RenderCleanTableWithBettingOptions(gameplay.SetBet, walletAmount)
	return
}

func playerDrawCard(gameplay *SinglePlayer) bool {
	card := gameplay.player.DrawCard(gameplay.deck)
	scores := gameplay.player.GetHandScores()
	gameplay.consoleUi.AddPlayerCard(card, scores)

	if gameplay.player.IsBusted() {
		gameplay.player.Loose()
		gameplay.whatsNext = nextStepNewGame
		gameplay.dealer.RevealSecondCard()
		gameplay.consoleUi.RenderDealerCards(nil)
		gameplay.consoleUi.RenderPlayerCards()
		gameplay.consoleUi.RenderPlayerBusted()
		gameplay.NewGame()
		return true
	}
	if gameplay.player.IsBlackjack() {
		gameplay.Stand()
		return true
	}
	return false
}

func dealerDrawCard(gameplay *SinglePlayer) bool {
	card := gameplay.dealer.DrawCard(gameplay.deck)
	scores := gameplay.dealer.GetHandScores()
	gameplay.consoleUi.AddDealerCard(card, scores)
	if gameplay.dealer.IsBusted() {
		gameplay.performPlayerWins()
		gameplay.whatsNext = nextStepNewGame
		gameplay.NewGame()
		return true
	}
	return false
}

func (gameplay *SinglePlayer) evaluate() {
	dealerBlackjack := gameplay.dealer.IsBlackjack()
	playerBlackjack := gameplay.player.IsBlackjack()

	if dealerBlackjack && playerBlackjack {
		gameplay.performDraw()
		return
	}
	if dealerBlackjack {
		gameplay.performDealerWins()
		return
	}
	if playerBlackjack {
		gameplay.performPlayerWins()
		return
	}
	if gameplay.player.GetHandScore() > gameplay.dealer.GetHandScore() {
		gameplay.performPlayerWins()
	} else if gameplay.player.GetHandScore() < gameplay.dealer.GetHandScore() {
		gameplay.performDealerWins()
	} else {
		gameplay.performDraw()
		return
	}
}

func (gameplay *SinglePlayer) performPlayerWins() {
	gameplay.player.Win()
	gameplay.dealer.RevealSecondCard()
	allDealerHandSum := gameplay.dealer.GetHandScores()
	gameplay.consoleUi.RenderDealerCards(allDealerHandSum)
	gameplay.consoleUi.RenderPlayerCards()
	gameplay.consoleUi.RenderPlayerWins()
}

func (gameplay *SinglePlayer) performDealerWins() {
	gameplay.dealer.Win()
	gameplay.player.Loose()
	gameplay.dealer.RevealSecondCard()
	allDealerHandSum := gameplay.dealer.GetHandScores()
	gameplay.consoleUi.RenderDealerCards(allDealerHandSum)
	gameplay.consoleUi.RenderPlayerCards()
	gameplay.consoleUi.RenderDealerWins()
}

func (gameplay *SinglePlayer) performDraw() {
	gameplay.player.Bet /= 2
	gameplay.player.Win()
	gameplay.dealer.Win()
	gameplay.dealer.RevealSecondCard()
	allDealerHandSum := gameplay.dealer.GetHandScores()
	gameplay.consoleUi.RenderDealerCards(allDealerHandSum)
	gameplay.consoleUi.RenderPlayerCards()
	gameplay.consoleUi.RenderDraw()
}
