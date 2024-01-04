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
			if got := tt.d.Add(tt.args.card); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
