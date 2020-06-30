package deck

import (
	"encoding/json"
	"strconv"
)

type CardType struct {
	name   string
	symbol rune
}

type SerializableCardType struct {
	Name   string
	Symbol rune
}

type Card struct {
	value     int
	cardType  *CardType
	IsVisible bool
}

type SerializableCard struct {
	Value     int
	CardType  *SerializableCardType
	IsVisible bool
}

func (c *Card) GetBlackjackValue() int {
	if c.value > 10 {
		return 10
	}
	return c.value
}

func (c *Card) GetSymbol() string {
	return string(c.cardType.symbol)
}

func (c *Card) GetDisplayingValue() string {
	switch c.value {
	case 1:
		return "A"
	case 12:
		return "J"
	case 13:
		return "Q"
	case 14:
		return "K"
	default:
		return strconv.Itoa(c.value)
	}
}

func (c *Card) Serialize() (result string, err error) {
	cardTypeJson := SerializableCardType{
		Name:   c.cardType.name,
		Symbol: c.cardType.symbol,
	}
	cardJson := SerializableCard{
		Value:     c.value,
		CardType:  &cardTypeJson,
		IsVisible: c.IsVisible,
	}
	b, err := json.Marshal(cardJson)
	if err != nil {
		return
	}
	result = string(b)
	return
}

func DeserializeCard(s string) Card {
	serializableCard := new(SerializableCard)
	b := []byte(s)
	json.Unmarshal(b, &serializableCard)
	return Card{
		value:     serializableCard.Value,
		cardType:  &CardType{
			name:   serializableCard.CardType.Name,
			symbol: serializableCard.CardType.Symbol,
		},
		IsVisible: serializableCard.IsVisible,
	}
}

//help with testing
func (c *Card) SetCard(value int, name string, symbol rune) {
	c.value = value
	c.cardType = &CardType{name, symbol}
}
