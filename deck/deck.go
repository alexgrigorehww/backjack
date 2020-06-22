package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type ShuffleType int

const (
	ShuffleAvailable = 1
	ShufflePast      = 2
	ShuffleAndMixAll = 3
)

type Deck struct {
	cards     []*Card
	discarded []*Card
}

func (d *Deck) Init() {
	d.cards = nil
	cardTypes := []CardType{{"clubs", '♣'}, {"diamonds", '♦'}, {"hearts", '♥'}, {"spades", '♠'}}
	for cardTypeIndex, _ := range cardTypes {
		for i := 1; i <= 14; i++ {
			if i == 11 {
				continue // 11 is ace which is 1
			}
			d.cards = append(d.cards, &Card{value: i, cardType: &cardTypes[cardTypeIndex]})
		}
	}
}

func (d *Deck) Shuffle(shuffleType ShuffleType){
	rand.Seed(time.Now().UnixNano())
	var whatToShuffle []*Card
	switch shuffleType {
	case ShuffleAvailable:
		whatToShuffle = d.cards
	case ShufflePast:
		whatToShuffle = d.discarded
	case ShuffleAndMixAll:
		d.cards = append(d.cards, d.discarded...)
		d.discarded = nil
		whatToShuffle = d.cards
	}
	rand.Shuffle(len(whatToShuffle), func(i, j int) {
		whatToShuffle[i], whatToShuffle[j] = whatToShuffle[j], whatToShuffle[i]
	})
}

func (d *Deck) Draw() *Card{
	if len(d.cards) == 0 {
		d.Shuffle(ShuffleAndMixAll)
	}
	var card *Card
	card, d.cards = d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]
	return card
}

func (d *Deck) Discard(cards []*Card){
	for _, card := range cards{
		card.IsVisible = false
	}
	d.discarded = append(d.discarded, cards...)
}

func (d *Deck) DiscardOne(card *Card){
	card.IsVisible = false
	d.discarded = append(d.discarded, card)
}

func (d *Deck) CardsLeft() int{
	return len(d.cards)
}

func (d *Deck) reveal() {
	for i, card := range d.cards {
		fmt.Printf("Card %d: \t %s \n", i, card.GetDisplayingValue()+card.GetSymbol())
	}
}
