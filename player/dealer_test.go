package player

import (
	"blackjack/deck"
	"testing"
)

func TestDealer_GetScore(t *testing.T) {
	mockedScore := 10
	score := (&Dealer{score: mockedScore}).GetScore()
	if score != mockedScore {
		t.Errorf("The dealer score should be %d", mockedScore)
	}
}

func TestDealer_Win(t *testing.T) {
	mockedScore := 10
	dealer := &Dealer{score: mockedScore}
	dealer.Win(2312)
	score := dealer.GetScore()
	if score != mockedScore + 1 {
		t.Errorf("The dealer score should be %d after winning", mockedScore + 1)
	}
}

func TestDealer_Loose(t *testing.T){
	mockedScore := 10
	dealer := &Dealer{score: mockedScore}
	dealer.Loose(2312)
	score := dealer.GetScore()
	if score != mockedScore - 1 {
		t.Errorf("The dealer score should be %d after winning", mockedScore - 1)
	}
}

func TestDealer_DrawCard(t *testing.T){
	var deck deck.Deck
	deck.Init()
	dealer := Dealer{}
	dealer.DrawCard(&deck)
	if len(dealer.hand.GetHandCards()) != 1{
		t.Error("Dealer should have 1 card")
	}
	dealer.DrawCard(&deck)
	if len(dealer.hand.GetHandCards()) != 2{
		t.Error("Dealer should have 2 cards")
	}
	dealer.DrawCard(&deck)
	if len(dealer.hand.GetHandCards()) != 3{
		t.Error("Dealer should have 3 cards")
	}
	card1 := dealer.hand.GetHandCards()[0]
	if !card1.IsVisible {
		t.Error("Dealer's first card should be visible")
	}
	card2 := dealer.hand.GetHandCards()[1]
	if card2.IsVisible {
		t.Error("Dealer's first card should be hidden")
	}
	card3 := dealer.hand.GetHandCards()[2]
	if !card3.IsVisible {
		t.Error("Dealer's first card should be visible")
	}
}