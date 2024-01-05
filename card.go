package xilac

import "bytes"

var (
	cardValueNames = []rune("A23456789_JQK")
	cardKindNames  = []rune("♥♦♣♠")
)

type CardRank uint8

func (r CardRank) Valid() bool {
	return r < 13
}

const (
	CardRankTwo CardRank = iota
	CardRankThree
	CardRankFour
	CardRankFive
	CardRankSix
	CardRankSeven
	CardRankEight
	CardRankNine
	CardRankTen
	CardRankJack
	CardRankQueen
	CardRankKing
	CardRankAce
)

type CardSuit uint8

func (s CardSuit) Valid() bool {
	return s < 4
}

const (
	CardSuitHeart CardSuit = iota
	CardSuitDiamond
	CardSuitClub
	CardSuitSpade
)

type Card uint8

func NewCard(rank CardRank, suit CardSuit) Card {
	if !rank.Valid() || !suit.Valid() {
		panic("invalid card")
	}

	return Card(uint8(suit)*13 + uint8(rank))
}

func NewCardFromID(id int) Card {
	return Card(id)
}

func (c Card) Valid() bool {
	return c < 52
}

func (c Card) Rank() CardRank {
	return CardRank(c % 13)
}

func (c Card) Suit() CardSuit {
	return CardSuit(c / 13)
}

func (c Card) String() string {
	v := c % 13
	bf := bytes.NewBuffer(nil)
	if v == 9 {
		bf.WriteString("10")
	} else {
		bf.WriteRune(cardValueNames[v])
	}
	bf.WriteRune(cardKindNames[c/13])
	return bf.String()
}
