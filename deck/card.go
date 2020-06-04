package deck

import "strconv"

type CardType struct {
	name   string
	symbol rune
}

type Card struct {
	value    int
	cardType *CardType
}

func (c *Card) GetBlackjackValue() int{
	if c.value > 10 {
		return c.value - 10
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