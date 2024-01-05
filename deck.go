package xilac

import (
	"crypto/rand"
	"math/big"
)

const (
	MaxCardNum = 52
	MaxPoint   = 21
)

type Deck []Card

type DeckStatus uint8

const (
	DeckStatusNormal DeckStatus = iota
	DeckStatusBusted
	DeckStatusBlackJack
	DeckStatusDoubleBlackJack
	DeckStatusHighFive
	DeckStatusTooHigh
	DeckStatusTooLow
)

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

func (d Deck) Value(minAcceptable, maxAcceptable int) (status DeckStatus, sum int) {
	if len(d) < 2 || len(d) > 5 {
		panic("invalid deck")
	}

	if len(d) == 2 {
		if d[0].IsAce() && d[1].IsAce() {
			return DeckStatusDoubleBlackJack, 0
		}
		if d[0].IsAce() && d[1].Value() == 10 || d[1].IsAce() && d[0].Value() == 10 {
			return DeckStatusBlackJack, 0
		}
	}

	sum, aceCnt := 0, 0
	for _, c := range d {
		if c.IsAce() {
			aceCnt++
		} else {
			sum += c.Value()
		}
	}

	if aceCnt > 0 {
		if sum >= MaxPoint-9 || len(d) >= 4 {
			sum += aceCnt
		} else if sum+11+(aceCnt-1) <= MaxPoint {
			sum += 11 + (aceCnt - 1)
		} else if sum+10+(aceCnt-1) <= MaxPoint {
			sum += 10 + (aceCnt - 1)
		} else {
			sum += aceCnt
		}
	}

	if sum > maxAcceptable {
		status = DeckStatusTooHigh
	} else if sum > MaxPoint {
		status = DeckStatusBusted
	} else if len(d) == 5 {
		status = DeckStatusHighFive
	} else if sum < minAcceptable {
		status = DeckStatusTooLow
	} else {
		status = DeckStatusNormal
	}

	return
}
