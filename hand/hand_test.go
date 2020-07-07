package hand_test

import (
	"blackjack/deck"
	"blackjack/hand"
	"testing"
)

func TestHand_AddCardToHand(t *testing.T) {
	var d deck.Deck
	d.Init()
	h := hand.Hand{}
	h.AddCardToHand(d.Draw())
	if len(h.GetHandCards()) == 0 {
		t.Error("No card added to hand. len=", len(h.GetHandCards()), '\n')
	}
}

func TestHand_GetHandCardsSum(t *testing.T) {
	var c1, c2, c3 deck.Card
	h := hand.Hand{}
	//case: bj
	c1.SetCard(10, "spades", '♠')
	c2.SetCard(1, "spades", '♠')
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c2)
	if h.GetHandCardsSum() != 21 {
		t.Error("This should be 21.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 2 aces + face card
	h.AddCardToHand(&c2)
	if h.GetHandCardsSum() != 12 {
		t.Error("This should be 12.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 3 aces
	h = hand.Hand{}
	c1.SetCard(1, "spades", '♠')
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c1)
	if h.GetHandCardsSum() != 13 {
		t.Error("This should be 13.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 2+9
	h = hand.Hand{}
	c1.SetCard(2, "spades", '♠')
	h.AddCardToHand(&c1)
	c2.SetCard(9, "spades", '♠')
	h.AddCardToHand(&c2)
	if h.GetHandCardsSum() != 11 {
		t.Error("This should be 11.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 3+A+6
	h = hand.Hand{}
	c1.SetCard(3, "spades", '♠')
	h.AddCardToHand(&c1)
	c2.SetCard(1, "spades", '♠')
	h.AddCardToHand(&c2)
	c3.SetCard(6, "spades", '♠')
	h.AddCardToHand(&c3)
	if h.GetHandCardsSum() != 20 {
		t.Error("This should be and 20.", h.GetHandCardsSum(), h.GetHandCards())
	}
}

func TestHand_DisplayValues(t *testing.T) {
	var c1, c2, c3 deck.Card
	h := hand.Hand{}

	//case: bj
	c1.SetCard(10, "spades", '♠')
	c1.IsVisible = true
	c2.SetCard(1, "spades", '♠')
	c2.IsVisible = true
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c2)
	if h.DisplayValues()[0] != 21 {
		t.Error("sMax should be 21.", h.DisplayValues())
	}

	//case: 2 aces + face card
	h.AddCardToHand(&c2)
	if h.DisplayValues()[0] != 12 {
		t.Error("sMin should be 12", h.DisplayValues())
	}

	//case: 3 aces
	h = hand.Hand{}
	c1.SetCard(1, "spades", '♠')
	c1.IsVisible = true
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c1)
	if h.DisplayValues()[0] != 3 {
		t.Error("This should be 3.", h.DisplayValues())
	}

	//case: 2 values returned
	c2.SetCard(7, "spades", '♠')
	c2.IsVisible = true
	h.AddCardToHand(&c2)
	if h.DisplayValues()[0] != 10 && h.DisplayValues()[1] != 20 {
		t.Error("sMion should be 10 and sMax should be 20", h.DisplayValues())
	}

	//case: 3 aces
	h = hand.Hand{}
	c1.SetCard(1, "spades", '♠')
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c1)
	h.AddCardToHand(&c1)
	if h.DisplayValues()[0] != 3 {
		t.Error("This should be 3.", h.DisplayValues())
	}
	//case: 2+3. At some point returned both sMin and sMax even though there are no aces
	h = hand.Hand{}
	c1.SetCard(2, "spades", '♠')
	c1.IsVisible = true
	h.AddCardToHand(&c1)
	c2.SetCard(3, "spades", '♠')
	c2.IsVisible = true
	h.AddCardToHand(&c2)
	if h.DisplayValues()[0] != 5 {
		t.Error("This should be 5.", h.DisplayValues())
	}
	//case: Hidden card
	h = hand.Hand{}
	c1.SetCard(2, "spades", '♠')
	h.AddCardToHand(&c1)
	c2.IsVisible = true
	c2.SetCard(3, "spades", '♠')
	c2.IsVisible = false
	h.AddCardToHand(&c2)
	if h.DisplayValues()[0] != 2 {
		t.Error("This should be 2.", h.DisplayValues())
	}
	//case: 2+9
	h = hand.Hand{}
	c1.SetCard(2, "spades", '♠')
	h.AddCardToHand(&c1)
	c2.SetCard(9, "spades", '♠')
	h.AddCardToHand(&c2)
	if h.DisplayValues()[0] == 11 {
		t.Error("This should be 11.", h.GetHandCardsSum(), h.GetHandCards())
	}
	//case: 3+A+6
	h = hand.Hand{}
	c1.SetCard(3, "spades", '♠')
	h.AddCardToHand(&c1)
	c2.SetCard(1, "spades", '♠')
	h.AddCardToHand(&c2)
	c3.SetCard(6, "spades", '♠')
	h.AddCardToHand(&c3)
	if h.DisplayValues()[0] == 10 && h.DisplayValues()[1] == 20 {
		t.Error("This should be 10 and 20.", h.GetHandCardsSum(), h.GetHandCards())
	}
}
