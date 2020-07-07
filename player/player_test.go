package player_test

import (
	"blackjack/deck"
	"blackjack/player"
	"testing"
)

func TestRegularPlayer_GetScore(t *testing.T) {
	score := (&player.RegularPlayer{}).GetScore()
	if score != 0 {
		t.Errorf("The regularPlayer score should be %d when is just created", 0)
	}
}

func TestRegularPlayer_Win(t *testing.T) {
	bet := 100
	var regularPlayer player.RegularPlayer
	regularPlayer.Init()
	initAmount := regularPlayer.GetWalletAmount()
	regularPlayer.Bet = bet
	regularPlayer.Win(bet)
	score := regularPlayer.GetScore()
	if score != 1 {
		t.Errorf("The regularPlayer score should be %d after winning", 1)
	}
	if regularPlayer.GetWalletAmount() != initAmount+bet {
		t.Errorf("The regularPlayer wallet amount should be %d after winning", bet)
	}
}

func TestRegularPlayer_Loose(t *testing.T) {
	bet := 10
	var regularPlayer player.RegularPlayer
	regularPlayer.Init()
	initAmount := regularPlayer.GetWalletAmount()
	regularPlayer.Bet = bet
	regularPlayer.Loose()
	score := regularPlayer.GetScore()
	if score != 0 {
		t.Errorf("The regularPlayer score should be %d after loosing", 0)
	}
	if regularPlayer.GetWalletAmount() != initAmount-bet {
		t.Errorf("The regularPlayer wallet amount should be %d after loosing", -bet)
	}
}

func TestRegularPlayer_DrawCard(t *testing.T) {
	var deck deck.Deck
	deck.Init()
	var regularPlayer player.RegularPlayer
	regularPlayer.Init()
	regularPlayer.DrawCard(&deck)
	if len(regularPlayer.GetCards()) != 1 {
		t.Error("regularPlayer should have 1 card")
	}
	regularPlayer.DrawCard(&deck)
	if len(regularPlayer.GetCards()) != 2 {
		t.Error("regularPlayer should have 2 cards")
	}
	regularPlayer.DrawCard(&deck)
	if len(regularPlayer.GetCards()) != 3 {
		t.Error("regularPlayer should have 3 cards")
	}
	if deck.CardsLeft() != 49 {
		t.Error("After 3 draws, deck should have 49 more cards")
	}
	card1 := regularPlayer.GetCards()[0]
	if !card1.IsVisible {
		t.Error("regularPlayer's first card should be visible")
	}
	card2 := regularPlayer.GetCards()[1]
	if !card2.IsVisible {
		t.Error("regularPlayer's second card should be hidden")
	}
	card3 := regularPlayer.GetCards()[2]
	if !card3.IsVisible {
		t.Error("regularPlayer's third card should be visible")
	}
}

func TestRegularPlayer_DiscardAllCards(t *testing.T) {
	var deck deck.Deck
	deck.Init()
	var regularPlayer player.RegularPlayer
	regularPlayer.Init()
	regularPlayer.DrawCard(&deck)
	regularPlayer.DrawCard(&deck)
	regularPlayer.DrawCard(&deck)
	regularPlayer.DiscardAllCards(&deck)
	if len(regularPlayer.GetCards()) != 0 {
		t.Error("After discarding all cards, in hand should be 0 cards")
	}
	if len(deck.GetDiscarded()) != 3 {
		t.Error("After discarding all cards, in hand should be 0 cards")
	}
}
