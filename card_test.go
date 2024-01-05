package xilac

import (
	"testing"
)

func newCardWithRank(rank CardRank) Card {
	return NewCard(rank, CardSuitHeart)
}

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		id   uint8
		want string
	}{
		{id: 0, want: "2♥"},
		{id: 8, want: "10♥"},
		{id: 38, want: "A♣"},
		{id: 50, want: "K♠"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card(tt.id)
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestCard_Valid(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want bool
	}{
		{c: Card(0), want: true},
		{c: Card(1), want: true},
		{c: Card(9), want: true},
		{c: Card(51), want: true},
		{c: Card(52), want: false},
		{c: Card(152), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSuit_Valid(t *testing.T) {
	tests := []struct {
		name string
		s    CardSuit
		want bool
	}{
		{s: CardSuitHeart, want: true},
		{s: CardSuitDiamond, want: true},
		{s: CardSuitClub, want: true},
		{s: CardSuitSpade, want: true},
		{s: CardSuit(4), want: false},
		{s: CardSuit(100), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardRank_Valid(t *testing.T) {
	tests := []struct {
		name string
		r    CardRank
		want bool
	}{
		{r: CardRankTwo, want: true},
		{r: CardRankThree, want: true},
		{r: CardRankAce, want: true},
		{r: CardRank(13), want: false},
		{r: CardRank(100), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCard(t *testing.T) {
	type args struct {
		rank CardRank
		suit CardSuit
	}
	tests := []struct {
		name string
		args args
		want Card
	}{
		{name: "A♥", args: args{rank: CardRankAce, suit: CardSuitHeart}, want: Card(12)},
		{name: "10♥", args: args{rank: CardRankTen, suit: CardSuitHeart}, want: Card(8)},
		{name: "K♣", args: args{rank: CardRankKing, suit: CardSuitClub}, want: Card(37)},
		{name: "K♠", args: args{rank: CardRankKing, suit: CardSuitSpade}, want: Card(50)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCard(tt.args.rank, tt.args.suit); got != tt.want {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Suit(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want CardSuit
	}{
		{name: "A♥", c: Card(12), want: CardSuitHeart},
		{name: "10♥", c: Card(8), want: CardSuitHeart},
		{name: "K♣", c: Card(37), want: CardSuitClub},
		{name: "K♠", c: Card(50), want: CardSuitSpade},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Suit(); got != tt.want {
				t.Errorf("Suit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Rank(t *testing.T) {
	tests := []struct {
		name string
		c    Card
		want CardRank
	}{
		{name: "A♥", c: Card(12), want: CardRankAce},
		{name: "10♥", c: Card(8), want: CardRankTen},
		{name: "K♣", c: Card(37), want: CardRankKing},
		{name: "K♠", c: Card(50), want: CardRankKing},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Rank(); got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}
