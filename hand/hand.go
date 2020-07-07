package hand

import (
	"blackjack/deck"
)

type Hand struct {
	cards []*deck.Card
}

type SerializableHand struct {
	Cards []*deck.SerializableCard
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
	if s+10+aces > 21 && aces != 0 { //s+11+aces-1>21
		s += aces
	} else if aces != 0 {
		s += 10 + aces //11-aces-1
	}
	return s
}

func (h *Hand) DisplayValues() []int {
	var s, aces int
	var scores []int
	aces = 0
	for i := 0; i < len(h.cards); i++ {
		if h.cards[i].IsVisible {
			//ignore aces
			if h.cards[i].GetBlackjackValue() == 1 {
				aces++ //treat aces separately
				continue
			}
			s += h.cards[i].GetBlackjackValue()
		}
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

func (h *Hand) GetSerializable() *SerializableHand {
	var serializableCards []*deck.SerializableCard
	for _, card := range h.cards {
		serializableCards = append(serializableCards, card.GetSerializable())
	}
	serializableHand := SerializableHand{
		Cards: serializableCards,
	}
	return &serializableHand
}

func (h *SerializableHand) Deserialize() *Hand {
	var cards []*deck.Card
	for _, card := range h.Cards {
		cards = append(cards, card.DeserializeCard())
	}
	hand := Hand{
		cards: cards,
	}
	return &hand
}
