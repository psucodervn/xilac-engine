package xilac

import "bytes"

const RankPerSuit = 13

var (
	cardValueNames = []rune("23456789_JQKA")
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

	return Card(uint8(suit)*RankPerSuit + uint8(rank))
}

func NewCardFromID(id int) Card {
	return Card(id)
}

func (c Card) Valid() bool {
	return c < 52
}

func (c Card) Rank() CardRank {
	return CardRank(c % RankPerSuit)
}

func (c Card) Suit() CardSuit {
	return CardSuit(c / RankPerSuit)
}

func (c Card) IsAce() bool {
	return c.Rank() == CardRankAce
}

func (c Card) Value() int {
	r := c.Rank()
	if r > CardRankTen {
		return 10
	}

	return int(r) + 2
}

func (c Card) String() string {
	r := c.Rank()
	bf := bytes.NewBuffer(nil)
	if r == CardRankTen {
		bf.WriteString("10")
	} else {
		bf.WriteRune(cardValueNames[r])
	}
	bf.WriteRune(cardKindNames[c.Suit()])
	return bf.String()
}
