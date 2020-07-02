package gameplay

import (
	"blackjack/deck"
	"blackjack/player"
	"blackjack/ui"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	ui        ui.UI
}

type SerializableSinglePlayer struct {
	WhatsNext string
	Dealer    *player.SerializableDealer
	Player    *player.SerializableRegularPlayer
	Deck      *deck.SerializableDeck
}

func (gameplay *SinglePlayer) Init(ui *ui.UI) (err error) {
	if gameplay.whatsNext != nextStepInit {
		err = errors.New("invalid gameplay state. cannot initialize the game")
		return
	}

	theDeck := new(deck.Deck)
	theDeck.Init()
	theDeck.Shuffle(deck.ShuffleAndMixAll)

	dealer := new(player.Dealer)
	dealer.Init()

	player := new(player.RegularPlayer)
	player.Init()

	// init UI
	gameplay.ui = *ui

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
	gameplay.ui.RenderDeal(gameplay.Deal)
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
	gameplay.ui.RenderHitOrStand(gameplay.Hit, gameplay.Stand)
	return
}

func (gameplay *SinglePlayer) Hit() (err error) {
	if gameplay.whatsNext != nextStepHitOrStand {
		err = errors.New("invalid gameplay state. you cannot hit")
		return
	}
	shouldStop := playerDrawCard(gameplay)
	if !shouldStop {
		gameplay.ui.RenderHitOrStand(gameplay.Hit, gameplay.Stand)
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
		gameplay.ui.RenderGameOver()
		return
	}
	gameplay.dealer.DiscardAllCards(gameplay.deck)
	gameplay.player.DiscardAllCards(gameplay.deck)
	gameplay.whatsNext = nextStepSetBet
	walletAmount := gameplay.player.GetWalletAmount()
	gameplay.ui.RenderCleanTableWithBettingOptions(gameplay.SetBet, gameplay.SaveGame, gameplay.SaveGame, walletAmount)
	return
}

func playerDrawCard(gameplay *SinglePlayer) bool {
	card := gameplay.player.DrawCard(gameplay.deck)
	scores := gameplay.player.GetHandScores()
	gameplay.ui.AddPlayerCard(card, scores)

	if gameplay.player.IsBusted() {
		gameplay.player.Loose()
		gameplay.whatsNext = nextStepNewGame
		gameplay.dealer.RevealSecondCard()
		allDealerHandSum := gameplay.dealer.GetHandScores()
		gameplay.ui.RenderDealerCards(allDealerHandSum)
		gameplay.ui.RenderPlayerCards()
		gameplay.ui.RenderPlayerBusted()
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
	gameplay.ui.AddDealerCard(card, scores)
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

func (gameplay *SinglePlayer) getSerializable() *SerializableSinglePlayer {
	serializableGameplay := SerializableSinglePlayer{
		WhatsNext: gameplay.whatsNext,
		Dealer:    gameplay.dealer.GetSerializable(),
		Player:    gameplay.player.GetSerializable(),
		Deck:      gameplay.deck.GetSerializable(),
	}
	return &serializableGameplay
}

func (gameplay *SinglePlayer) SaveGame() error {
	serializable := gameplay.getSerializable()
	b, err := json.Marshal(serializable)
	if err != nil {
		return err
	}
	ioutil.WriteFile("restoreFile.json", b, 0644)
	return nil
}

func (gameplay *SinglePlayer) performPlayerWins() {
	gameplay.player.Win()
	gameplay.dealer.RevealSecondCard()
	allDealerHandSum := gameplay.dealer.GetHandScores()
	gameplay.ui.RenderDealerCards(allDealerHandSum)
	gameplay.ui.RenderPlayerCards()
	gameplay.ui.RenderPlayerWins()
}

func (gameplay *SinglePlayer) performDealerWins() {
	gameplay.dealer.Win()
	gameplay.player.Loose()
	gameplay.dealer.RevealSecondCard()
	allDealerHandSum := gameplay.dealer.GetHandScores()
	gameplay.ui.RenderDealerCards(allDealerHandSum)
	gameplay.ui.RenderPlayerCards()
	gameplay.ui.RenderDealerWins()
}

func (gameplay *SinglePlayer) performDraw() {
	gameplay.player.Bet /= 2
	gameplay.player.Win()
	gameplay.dealer.Win()
	gameplay.dealer.RevealSecondCard()
	allDealerHandSum := gameplay.dealer.GetHandScores()
	gameplay.ui.RenderDealerCards(allDealerHandSum)
	gameplay.ui.RenderPlayerCards()
	gameplay.ui.RenderDraw()
}
