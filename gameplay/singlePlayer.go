package gameplay

import (
	"blackjack/deck"
	"blackjack/player"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const(
	nextStepInit   = ""
	nextStepSetBet = "SET_BET"
	nextStepDeal   = "DEAL"
	nextStepHitOrStand = "HIT_OR_STAND"
	nextStepNewGame = "NEW_GAME"
)

type SinglePlayer struct {
	whatsNext string
	dealer *player.Dealer
	player *player.RegularPlayer
	deck *deck.Deck
}

func (gameplay *SinglePlayer) Init() (err error){
	if gameplay.whatsNext != nextStepInit {
		err = errors.New("invalid gameplay state. cannot initialize the game")
		return
	}
	theDeck := new(deck.Deck)
	theDeck.Init()
	theDeck.Shuffle(deck.ShuffleAndMixAll)

	dealer := new(player.Dealer)
	dealer.Init()

	// for the moment only one player
	// in future we may add players from UI (join style)
	player := new(player.RegularPlayer)
	player.Init()

	// init gameplay
	gameplay.deck = theDeck
	gameplay.dealer = dealer
	gameplay.player = player

	gameplay.whatsNext = nextStepSetBet
	renderCleanTableWithBettingOptions(gameplay)
	return
}

func (gameplay *SinglePlayer) SetBet(bet int) (err error){
	if gameplay.whatsNext != nextStepSetBet {
		err = errors.New("invalid gameplay state. you cannot set bet")
		return
	}
	if gameplay.player.GetWalletAmount() < bet {
		err = errors.New("insufficient amount")
		return
	}
	gameplay.player.Bet = bet
	renderDeal(gameplay)
	gameplay.whatsNext = nextStepDeal
	return
}

func (gameplay *SinglePlayer) Deal() (err error){
	if gameplay.whatsNext != nextStepSetBet {
		err = errors.New("invalid gameplay state. you cannot deal")
		return
	}
	playerDrawCard(gameplay)
	dealerDrawCard(gameplay)
	playerDrawCard(gameplay)
	dealerDrawCard(gameplay)

	gameplay.whatsNext = nextStepHitOrStand
	renderHitOrStand(gameplay)
	return
}

func (gameplay *SinglePlayer) Hit() (err error){
	if gameplay.whatsNext != nextStepHitOrStand {
		err = errors.New("invalid gameplay state. you cannot hit")
		return
	}
	shouldStop := playerDrawCard(gameplay)

	if !shouldStop {
		renderHitOrStand(gameplay)
	}
	return
}

func (gameplay *SinglePlayer) Stand() (err error){
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

func (gameplay *SinglePlayer) NewGame() (err error){
	if gameplay.whatsNext != nextStepNewGame {
		err = errors.New("invalid gameplay state. you cannot start new game")
		return
	}
	gameplay.dealer.DiscardAllCards(gameplay.deck)
	gameplay.player.DiscardAllCards(gameplay.deck)
	gameplay.whatsNext = nextStepSetBet
	renderCleanTableWithBettingOptions(gameplay)
	return
}

func playerDrawCard(gameplay *SinglePlayer) bool{
	card := gameplay.player.DrawCard(gameplay.deck)
	scores := gameplay.player.GetHandScores()
	renderPlayerCardAdded(card, scores)

	if gameplay.player.IsBusted(){
		gameplay.player.Loose()
		gameplay.whatsNext = nextStepNewGame
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderPlayerBusted()
		gameplay.NewGame()
		return true
	}
	if gameplay.player.IsBlackjack(){
		gameplay.Stand()
		return true
	}
	return false
}

func dealerDrawCard(gameplay *SinglePlayer) bool{
	card := gameplay.dealer.DrawCard(gameplay.deck)
	scores := gameplay.dealer.GetHandScores()
	renderDealerCardAdded(card, scores)
	if gameplay.dealer.IsBusted(){
		gameplay.player.Win()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderPlayerWins()
		gameplay.whatsNext = nextStepNewGame
		gameplay.NewGame()
		return true
	}
	return false
}

func (gameplay *SinglePlayer) evaluate() {
	dealerBlackjack := gameplay.dealer.IsBlackjack()
	playerBlackjack := gameplay.player.IsBlackjack()

	if dealerBlackjack && playerBlackjack{
		gameplay.player.Bet /= 2
		gameplay.player.Win()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderDraw()
		return
	}
	if dealerBlackjack{
		gameplay.dealer.Win()
		gameplay.player.Loose()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderDealerWins()
		return
	}
	if playerBlackjack{
		gameplay.player.Win()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderPlayerWins()
		return
	}
	println(gameplay.player.GetHandScore())
	println(gameplay.dealer.GetHandScore())
	if gameplay.player.GetHandScore() > gameplay.dealer.GetHandScore() {
		gameplay.player.Win()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderPlayerWins()
	} else if gameplay.player.GetHandScore() < gameplay.dealer.GetHandScore() {
		gameplay.dealer.Win()
		gameplay.player.Loose()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderDealerWins()
	} else {
		gameplay.player.Bet /= 2
		gameplay.player.Win()
		renderCards(gameplay.dealer)
		renderCards(gameplay.player)
		renderDraw()
		return
	}
}
// for rendering
func renderCleanTableWithBettingOptions(gameplay *SinglePlayer){
	bet, err := strconv.Atoi(read("New Game! Your wallet: "+ strconv.Itoa(gameplay.player.GetWalletAmount()) +". Type your bet"))
	if err!=nil {
		renderCleanTableWithBettingOptions(gameplay)
	} else {
		err := gameplay.SetBet(bet)
		if err != nil{
			fmt.Println(err)
		}
	}
}

func renderDeal(gameplay *SinglePlayer){
	deal := read("You want to deal? (y/n)")
	if deal == "y"{
		gameplay.Deal()
	} else {
		renderDeal(gameplay)
	}
}

func renderHitOrStand(gameplay *SinglePlayer){
	renderCards(gameplay.dealer)
	renderCards(gameplay.player)
	action := read("HIT (h) / STAND (s)")
	if action == "h" {
		err := gameplay.Hit()
		if err != nil {
			fmt.Println(err)
		}
	} else if action == "s" {
		err := gameplay.Stand()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		renderHitOrStand(gameplay)
	}
}

func renderPlayerCardAdded(card *deck.Card, playerSums []int){
	// todo: render card added for player
	// todo: render hand sums for player
}

func renderDealerCardAdded(card *deck.Card, dealerSums []int){
	// todo: render card added for dealer
	// todo: render hand sums for dealer
}

func renderPlayerBusted(){
	println("Player busted")
	// todo: render busted
}

func renderPlayerWins(){
	println("Player wins!")
	// todo: render player wins
}

func renderDraw(){
	println("DRAW")
	// todo: render draw
}

func renderDealerWins(){
	println("Dealer wins")
	// todo: render dealer wins
}

func read(label string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label + ": ")
	text, _ := reader.ReadString('\n')
	return strings.Trim(text, "\n\r")
}

// not needed in real renderer
func renderCards(player player.Player){
	cards := player.GetCards()
	displayingValues := ""
	for _, card := range cards {
		if !card.IsVisible {
			displayingValues += " ? "
		} else {
			displayingValues += " " + card.GetDisplayingValue() + " "
		}
	}
	fmt.Printf("%T cards: %s Sum: %d \n", player, displayingValues, player.GetHandScores())
}