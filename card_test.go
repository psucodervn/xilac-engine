package xilac

import (
	"testing"
)

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		id   uint8
		want string
	}{
		{id: 0, want: "A♥"},
		{id: 9, want: "10♥"},
		{id: 38, want: "K♣"},
		{id: 51, want: "K♠"},
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

func TestCard_Value(t *testing.T) {
	tests := []struct {
		name string
		id   uint8
		want uint8
	}{
		{id: 0, want: 1},
		{id: 1, want: 2},
		{id: 9, want: 10},
		{id: 10, want: 10},
		{id: 51, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card(tt.id)
			if got := c.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
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

func TestCard_Compare(t *testing.T) {
	type args struct {
		other Card
	}
	tests := []struct {
		name string
		c    Card
		args args
		want int
	}{
		{c: NewCard(CardRankAce, CardSuitClub), args: args{other: NewCard(CardRankAce, CardSuitClub)}, want: 0},
		{c: NewCard(CardRankAce, CardSuitClub), args: args{other: NewCard(CardRankAce, CardSuitDiamond)}, want: 0},
		{c: NewCard(CardRankAce, CardSuitClub), args: args{other: NewCard(CardRankTwo, CardSuitClub)}, want: 1},
		{c: NewCard(CardRankJack, CardSuitClub), args: args{other: NewCard(CardRankKing, CardSuitDiamond)}, want: 0},
		{c: NewCard(CardRankNine, CardSuitClub), args: args{other: NewCard(CardRankAce, CardSuitClub)}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Compare(tt.args.other); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
