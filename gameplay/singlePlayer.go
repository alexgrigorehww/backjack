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
	restoreFile        = "restoreFile.json"
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
	gameplay.newGame()

	return
}

func (gameplay *SinglePlayer) setBet(bet int) (err error) {
	if gameplay.whatsNext != nextStepSetBet {
		err = errors.New("invalid gameplay state. you cannot set bet")
		return
	}
	if gameplay.player.GetWalletAmount() < bet {
		err = errors.New("insufficient amount")
		return
	}
	gameplay.player.Bet = bet
	gameplay.ui.RenderDeal(gameplay.deal)
	gameplay.whatsNext = nextStepDeal
	return
}

func (gameplay *SinglePlayer) deal() (err error) {
	if gameplay.whatsNext != nextStepSetBet {
		err = errors.New("invalid gameplay state. you cannot deal")
		return
	}
	playerDrawCard(gameplay)
	dealerDrawCard(gameplay)
	playerDrawCard(gameplay)
	dealerDrawCard(gameplay)

	gameplay.whatsNext = nextStepHitOrStand
	gameplay.ui.RenderHitOrStand(gameplay.hit, gameplay.stand)
	return
}

func (gameplay *SinglePlayer) hit() (err error) {
	if gameplay.whatsNext != nextStepHitOrStand {
		err = errors.New("invalid gameplay state. you cannot hit")
		return
	}
	shouldStop := playerDrawCard(gameplay)
	if !shouldStop {
		gameplay.ui.RenderHitOrStand(gameplay.hit, gameplay.stand)
	}
	return
}

func (gameplay *SinglePlayer) stand() (err error) {
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
			gameplay.newGame()
			break
		}
	}
	return
}

func (gameplay *SinglePlayer) newGame() (err error) {
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
	gameplay.ui.RenderCleanTableWithBettingOptions(gameplay.setBet, gameplay.saveGame, gameplay.restoreGame, walletAmount)
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
		gameplay.newGame()
		return true
	}
	if gameplay.player.IsBlackjack() {
		gameplay.stand()
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
		gameplay.newGame()
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

func (serializableGameplay *SerializableSinglePlayer) deserialize() *SinglePlayer {
	gameplay := SinglePlayer{
		whatsNext: serializableGameplay.WhatsNext,
		dealer:    serializableGameplay.Dealer.Deserialize(),
		player:    serializableGameplay.Player.Deserialize(),
		deck:      serializableGameplay.Deck.DeserializeDeck(),
	}
	return &gameplay
}

func (gameplay *SinglePlayer) saveGame(ch chan error) {
	serializable := gameplay.getSerializable()
	b, err := json.Marshal(serializable)
	if err != nil {
		ch <- err
		return
	}
	ioutil.WriteFile(restoreFile, b, 0644)
	gameplay.whatsNext = nextStepNewGame
	gameplay.newGame()
	ch <- nil
}

func (gameplay *SinglePlayer) restoreGame(ch chan error) {
	b, err := ioutil.ReadFile(restoreFile)
	if err != nil {
		ch <- err
		return
	}
	var serializableSinglePlayer SerializableSinglePlayer
	err = json.Unmarshal(b, &serializableSinglePlayer)
	if err != nil {
		ch <- err
		return
	}
	restoredGameplay := serializableSinglePlayer.deserialize()
	gameplay.deck = restoredGameplay.deck
	gameplay.player = restoredGameplay.player
	gameplay.dealer = restoredGameplay.dealer
	gameplay.whatsNext = restoredGameplay.whatsNext

	gameplay.whatsNext = nextStepNewGame
	gameplay.newGame()
	ch <- nil
	return
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
