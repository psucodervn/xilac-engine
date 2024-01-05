package xilac

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	type args struct {
		cards []Card
	}
	tests := []struct {
		name string
		args args
		want Deck
	}{
		{name: "empty", args: args{cards: []Card{}}, want: Deck{}},
		{
			name: "two cards",
			args: args{cards: []Card{4, 5}},
			want: Deck{Card(4), Card(5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(tt.args.cards...); !got.Equals(tt.want) {
				t.Errorf("NewDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Equals(t *testing.T) {
	type args struct {
		other Deck
	}
	tests := []struct {
		name string
		d    Deck
		args args
		want bool
	}{
		{name: "empty", d: Deck{}, args: args{other: Deck{}}, want: true},
		{
			name: "two cards",
			d:    Deck{Card(4), Card(5)},
			args: args{other: Deck{Card(4), Card(5)}},
			want: true,
		},
		{
			name: "two cards",
			d:    Deck{Card(4), Card(6)},
			args: args{other: Deck{Card(4), Card(5)}},
			want: false,
		},
		{
			name: "two cards with one invalid",
			d:    Deck{Card(4)},
			args: args{other: Deck{Card(100), Card(4)}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Add(t *testing.T) {
	type args struct {
		card Card
	}
	tests := []struct {
		name string
		d    Deck
		args args
		want Deck
	}{
		{name: "empty", d: Deck{}, args: args{card: Card(4)}, want: Deck{Card(4)}},
		{
			name: "two cards",
			d:    Deck{Card(4), Card(5)},
			args: args{card: Card(6)},
			want: Deck{Card(4), Card(5), Card(6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Push(tt.args.card); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Value(t *testing.T) {
	tests := []struct {
		name    string
		d       Deck
		want    DeckStatus
		wantSum int
	}{
		{
			name:    "too low",
			d:       Deck{newCardWithRank(CardRankTwo), newCardWithRank(CardRankThree)},
			want:    DeckStatusTooLow,
			wantSum: 5,
		},
		{
			name:    "normal",
			d:       Deck{newCardWithRank(CardRankFive), newCardWithRank(CardRankAce)},
			want:    DeckStatusNormal,
			wantSum: 16,
		},
		{
			name:    "blackjack",
			d:       Deck{newCardWithRank(CardRankJack), newCardWithRank(CardRankAce)},
			want:    DeckStatusBlackJack,
			wantSum: 0,
		},
		{
			name:    "double blackjack",
			d:       Deck{newCardWithRank(CardRankAce), newCardWithRank(CardRankAce)},
			want:    DeckStatusDoubleBlackJack,
			wantSum: 0,
		},
		{
			name:    "busted",
			d:       Deck{newCardWithRank(CardRankFive), newCardWithRank(CardRankSeven), newCardWithRank(CardRankTen)},
			want:    DeckStatusBusted,
			wantSum: 22,
		},
		{
			name:    "too high",
			d:       Deck{newCardWithRank(CardRankTen), newCardWithRank(CardRankNine), newCardWithRank(CardRankTen)},
			want:    DeckStatusTooHigh,
			wantSum: 29,
		},
		{
			name:    "high five",
			d:       Deck{newCardWithRank(CardRankTwo), newCardWithRank(CardRankThree), newCardWithRank(CardRankFour), newCardWithRank(CardRankFive), newCardWithRank(CardRankTwo)},
			want:    DeckStatusHighFive,
			wantSum: 16,
		},
		{
			name:    "high five",
			d:       Deck{newCardWithRank(CardRankTwo), newCardWithRank(CardRankThree), newCardWithRank(CardRankFour), newCardWithRank(CardRankAce), newCardWithRank(CardRankTwo)},
			want:    DeckStatusHighFive,
			wantSum: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, sum := tt.d.Value(16, 27)
			if status != tt.want || sum != tt.wantSum {
				t.Errorf("Value() = %v, %v, want %v, %v", status, sum, tt.want, tt.wantSum)
			}
		})
	}
}
