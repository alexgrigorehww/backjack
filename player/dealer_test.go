package player_test

import (
	"blackjack/deck"
	"blackjack/player"
	"testing"
)

func TestDealer_GetScore(t *testing.T) {
	var dealer player.Dealer
	dealer.Init()
	score := dealer.GetScore()
	if score != 0 {
		t.Errorf("The dealer score should be %d when is just created", 0)
	}
}

func TestDealer_Win(t *testing.T) {
	var dealer player.Dealer
	dealer.Init()
	dealer.Win()
	score := dealer.GetScore()
	if score != 1 {
		t.Errorf("The dealer score should be %d after winning", 1)
	}
}

func TestDealer_Loose(t *testing.T) {
	var dealer player.Dealer
	dealer.Init()
	dealer.Loose()
	score := dealer.GetScore()
	if score != -1 {
		t.Errorf("The dealer score should be %d after winning", -1)
	}
}

func TestDealer_DrawCard(t *testing.T) {
	var deck deck.Deck
	deck.Init()
	var dealer player.Dealer
	dealer.Init()
	dealer.DrawCard(&deck)
	if len(dealer.GetCards()) != 1 {
		t.Error("Dealer should have 1 card")
	}
	dealer.DrawCard(&deck)
	if len(dealer.GetCards()) != 2 {
		t.Error("Dealer should have 2 cards")
	}
	dealer.DrawCard(&deck)
	if len(dealer.GetCards()) != 3 {
		t.Error("Dealer should have 3 cards")
	}
	if deck.CardsLeft() != 49 {
		t.Error("After 3 draws, deck should have 49 more cards")
	}
	card1 := dealer.GetCards()[0]
	if !card1.IsVisible {
		t.Error("Dealer's first card should be visible")
	}
	card2 := dealer.GetCards()[1]
	if card2.IsVisible {
		t.Error("Dealer's second card should be hidden")
	}
	card3 := dealer.GetCards()[2]
	if !card3.IsVisible {
		t.Error("Dealer's third card should be visible")
	}
}

func TestDealer_DiscardAllCards(t *testing.T) {
	var deck deck.Deck
	deck.Init()
	var dealer player.Dealer
	dealer.Init()
	dealer.DrawCard(&deck)
	dealer.DrawCard(&deck)
	dealer.DrawCard(&deck)
	dealer.DiscardAllCards(&deck)
	if len(dealer.GetCards()) != 0 {
		t.Error("After discarding all cards, in hand should be 0 cards")
	}
	if len(deck.GetDiscarded()) != 3 {
		t.Error("After discarding all cards, in hand should be 0 cards")
	}
}
