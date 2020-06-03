package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	nextCardIndex int
	cards [52]Card
}

func (d *Deck) Init() {
	cardTypes := []CardType{{"clubs", '♣'}, {"diamonds", '♦'}, {"hearts", '♥'}, {"spades", '♠'}}
	tmpCardIndex := 0
	for cardTypeIndex, _ := range cardTypes {
		tmpCardIndex = 0
		for i := 1; i <= 14; i++ {
			if i == 11 {
				continue // 11 is ace which is 1
			}
			index := cardTypeIndex*13 + tmpCardIndex
			tmpCardIndex++
			d.cards[index] = Card{i, &cardTypes[cardTypeIndex]}
		}
	}
}

func (d *Deck) Shuffle(){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() Card{
	if d.nextCardIndex > 51 {
		d.nextCardIndex = 0
	}
	card := d.cards[d.nextCardIndex]
	d.nextCardIndex++
	return card
}

func (d *Deck) CardsLeft() int{
	return 52 - d.nextCardIndex
}

func (d *Deck) reveal() {
	for i, card := range d.cards {
		fmt.Printf("Card %d: \t %s \n", i, card.GetDisplayingValue()+card.GetSymbol())
	}
}
