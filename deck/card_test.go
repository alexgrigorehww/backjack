package deck

import (
	"testing"
)

func TestCard_GetDisplayingValue(t *testing.T) {
	if (&Card{value: 10}).GetDisplayingValue() != "10" {
		t.Error("10 should be string 10")
	}

	if (&Card{value: 12}).GetDisplayingValue() != "J" {
		t.Error("12 should be J")
	}

	if (&Card{value: 13}).GetDisplayingValue() != "Q" {
		t.Error("13 should be Q")
	}

	if (&Card{value: 14}).GetDisplayingValue() != "K" {
		t.Error("14 should be K")
	}

	if (&Card{value: 1}).GetDisplayingValue() != "A" {
		t.Error("1 should be A")
	}
}

func TestCard_GetSymbol(t *testing.T) {
	cardType := CardType{symbol: 'S'}
	if (&Card{cardType: &cardType}).GetSymbol() != "S" {
		t.Error("Card symbol S should be displayed as string S")
	}
}

func TestCard_GetBlackjackValue(t *testing.T) {
	if (&Card{value: 10}).GetBlackjackValue() != 10 {
		t.Error("10 should be 10")
	}
	if (&Card{value: 5}).GetBlackjackValue() != 5 {
		t.Error("5 should be 5")
	}
	if (&Card{value: 14}).GetBlackjackValue() != 4 {
		t.Error("14 should be 4")
	}
}