package hand

import (
	"blackjack/deck"
	"testing"
)

func TestHand_AddCardToHand(t *testing.T) {
	var d deck.Deck
	d.Init()
	h := Hand{}
	h.AddCardToHand(d.Draw())
	if len(h.cards) == 0 {
		t.Error("No card added to hand. len=", len(h.cards), '\n')
	}
}

func TestHand_GetHandCardsSum(t *testing.T) {
	var c1, c2 deck.Card
	h := Hand{}
	//case: bj
	c1.SetCard(10, "spades", '♠')
	c2.SetCard(1, "spades", '♠')
	h.AddCardToHand(c1)
	h.AddCardToHand(c2)
	if h.GetHandCardsSum() != 21 {
		t.Error("This should be 21.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 2 aces + face card
	h.AddCardToHand(c2)
	if h.GetHandCardsSum() != 12 {
		t.Error("This should be 12.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 3 aces
	h = Hand{}
	c1.SetCard(1, "spades", '♠')
	h.AddCardToHand(c1)
	h.AddCardToHand(c1)
	h.AddCardToHand(c1)
	if h.GetHandCardsSum() != 13 {
		t.Error("This should be 13.", h.GetHandCardsSum(), h.GetHandCards())
	}
}

func TestHand_DisplayValues(t *testing.T) {
	var c1, c2 deck.Card
	h := Hand{}

	//case: bj
	c1.SetCard(10, "spades", '♠')
	c2.SetCard(1, "spades", '♠')
	h.AddCardToHand(c1)
	h.AddCardToHand(c2)
	if h.DisplayValues()[0] != 21 {
		t.Error("sMax should be 21.", h.DisplayValues())
	}

	//case: 2 aces + face card
	h.AddCardToHand(c2)
	if h.DisplayValues()[0] != 12 {
		t.Error("sMin should be 12", h.DisplayValues())
	}

	//case: 3 aces
	h = Hand{}
	c1.SetCard(1, "spades", '♠')
	h.AddCardToHand(c1)
	h.AddCardToHand(c1)
	h.AddCardToHand(c1)
	if h.DisplayValues()[0] != 3 {
		t.Error("This should be 3.", h.DisplayValues())
	}

	//case: 2 values returned
	c2.SetCard(7, "spades", '♠')
	h.AddCardToHand(c2)
	if h.DisplayValues()[0] != 10 && h.DisplayValues()[1] != 20 {
		t.Error("sMion should be 10 and sMax should be 20", h.DisplayValues())
	}
}
