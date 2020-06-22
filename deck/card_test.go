package deck_test

import (
	"blackjack/deck"
	"testing"
)

func TestCard_GetDisplayingValue(t *testing.T) {
	var card1, card2, card3, card4, card5 deck.Card
	card1.SetCard(10, "", ' ')
	if card1.GetDisplayingValue() != "10" {
		t.Error("10 should be string 10")
	}
	card2.SetCard(12, "", ' ')
	if card2.GetDisplayingValue() != "J" {
		t.Error("12 should be J")
	}
	card3.SetCard(13, "", ' ')
	if card3.GetDisplayingValue() != "Q" {
		t.Error("13 should be Q")
	}
	card4.SetCard(14, "", ' ')
	if card4.GetDisplayingValue() != "K" {
		t.Error("14 should be K")
	}
	card5.SetCard(1, "", ' ')
	if card5.GetDisplayingValue() != "A" {
		t.Error("1 should be A")
	}
}

func TestCard_GetSymbol(t *testing.T) {
	var card deck.Card
	card.SetCard(1,"", 'S')
	if card.GetSymbol() != "S" {
		t.Error("Card symbol S should be displayed as string S")
	}
}

func TestCard_GetBlackjackValue(t *testing.T) {
	var card1, card2, card3 deck.Card
	card1.SetCard(10, "", ' ')
	if card1.GetBlackjackValue() != 10 {
		t.Error("10 should be 10")
	}
	card2.SetCard(5, "", ' ')
	if card2.GetBlackjackValue() != 5 {
		t.Error("5 should be 5")
	}
	card3.SetCard(14, "", ' ')
	if card3.GetBlackjackValue() != 10 {
		t.Error("14 should be 10")
	}
}