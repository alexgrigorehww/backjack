package hand

import (
	"blackjack/deck"
)

type Hand struct {
	cards []*deck.Card
}

//Add card to hand
func (h *Hand) AddCardToHand(c *deck.Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) DiscardAllCards() {
	h.cards = nil
}

func (h *Hand) GetHandCards() []*deck.Card {
	return h.cards
}

func (h *Hand) GetHandCardsSum() int {
	var s, aces int
	aces = 0
	for i := 0; i < len(h.cards); i++ {
		//ignore aces
		if h.cards[i].GetBlackjackValue() == 1 {
			aces++ //treat aces separately
			continue
		}
		s += h.cards[i].GetBlackjackValue()
	}
	if s+10+aces > 21 { //s+11+aces-1>21
		s += aces
	} else {
		s += 10 + aces //11-aces-1
	}
	return s
}

func (h *Hand) DisplayValues() []int {
	var s, aces int
	var scores []int
	aces = 0
	for i := 0; i < len(h.cards); i++ {
		//ignore aces
		if h.cards[i].GetBlackjackValue() == 1 {
			aces++ //treat aces separately
			continue
		}
		s += h.cards[i].GetBlackjackValue()
	}
	//sMin=s+aces
	//sMax=s+aces+10
	switch {
	case s+aces+10 == 21 || s+aces == 21:
		scores = append(scores, 21)
	case s+aces+10 < 21 && s+aces < 21 && aces != 0:
		scores = append(scores, s+aces, s+aces+10)
	default:
		scores = append(scores, s+aces)
	}
	return scores
}
