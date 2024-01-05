package xilac

import (
	"crypto/rand"
	"math/big"
)

const MaxCardNum = 52

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

func NewShuffledFullDeck() Deck {
	deck := make(Deck, MaxCardNum)
	for i := range deck {
		deck[i] = NewCardFromID(i)
	}

	// shuffle cards
	for i := MaxCardNum - 1; i > 0; i-- {
		bj, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
		j := bj.Int64()
		deck[i], deck[j] = deck[j], deck[i]
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

func (d Deck) Push(card Card) Deck {
	if !card.Valid() {
		panic("invalid card")
	}
	return append(d, card)
}

func (d Deck) Pop() (Card, Deck) {
	if len(d) == 0 {
		panic("empty deck")
	}
	return d[len(d)-1], d[:len(d)-1]
}
