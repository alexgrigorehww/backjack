package deck

import (
	"testing"
)

func checkExpectedCard(i int, card Card) (expectedValue int, ok bool) {
	ok = true
	expectedValue = i%13 + 1
	if expectedValue > 10{
		expectedValue++
	}
	if card.value != expectedValue {
		ok = false
	}
	return
}

func TestDeck_Init(t *testing.T) {
	var deck Deck
	deck.Init()
	// check cards value
	for i, card := range deck.cards {
		expectedValue, ok := checkExpectedCard(i, *card)
		if !ok {
			t.Errorf("The %d card should be %d", i, expectedValue)
		}
	}
	// check the cards type
	for i, card := range deck.cards {
		switch {
		case i<13:
			if card.cardType.symbol != '♣' || card.cardType.name != "clubs"{
				t.Errorf("The %d card should be clubs", i)
			}
		case i<26:
			if card.cardType.symbol != '♦' || card.cardType.name != "diamonds"{
				t.Errorf("The %d card should be diamonds", i)
			}
		case i<39:
			if card.cardType.symbol != '♥' || card.cardType.name != "hearts"{
				t.Errorf("The %d card should be hearts", i)
			}
		default:
			if card.cardType.symbol != '♠' || card.cardType.name != "spades"{
				t.Errorf("The %d card should be spades", i)
			}
		}
	}
}

func TestDeck_Shuffle(t *testing.T){
	var deck Deck
	deck.Init()
	deck.Shuffle(ShuffleAndMixAll)
	if deck.cards[0].value == 1 && deck.cards[1].value == 2{
		t.Error("Shuffle seems odd")
	}
}

func TestDeck_Draw(t *testing.T) {
	var deck Deck
	deck.Init()

	for left := deck.CardsLeft(); left > 0; left = deck.CardsLeft(){
		card := deck.Draw()
		cardIndex := left-1
		expectedValue, ok:=checkExpectedCard(cardIndex, *card)
		if !ok {
			t.Errorf("Draw has an unexpected behaviuor: The %d card should be %d, but it's %d", cardIndex, expectedValue, card.value)
		}
	}
}