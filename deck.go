package xilac

type Deck []Card

func NewDeck(cards ...Card) Deck {
	deck := make(Deck, 0, len(cards))
	for _, card := range cards {
		if !card.Valid() {
			panic("invalid card")
		}
		deck = append(deck, card)
	}
	return deck
}

func (d Deck) Equals(other Deck) bool {
	if len(d) != len(other) {
		return false
	}
	for i, c := range d {
		if c != other[i] {
			return false
		}
	}
	return true
}

func (d Deck) Add(card Card) Deck {
	if !card.Valid() {
		panic("invalid card")
	}
	return append(d, card)
}
