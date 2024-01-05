package xilac

import "testing"

func TestGame_deal(t *testing.T) {
	type fields struct {
		players []GamePlayer
		dealer  GamePlayer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "deal",
			fields: fields{
				players: []GamePlayer{
					{
						Deck:   Deck{},
						Status: GamePlayerStatusWaiting,
					},
					{
						Deck:   Deck{},
						Status: GamePlayerStatusWaiting,
					},
				},
				dealer: GamePlayer{
					Deck:   Deck{},
					Status: GamePlayerStatusWaiting,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				players: tt.fields.players,
				dealer:  tt.fields.dealer,
			}
			g.deal()
			for _, p := range g.players {
				if len(p.Deck) != 2 {
					t.Errorf("deal() = %v, want %v", len(p.Deck), 2)
				}
			}
			if len(g.dealer.Deck) != 2 {
				t.Errorf("deal() = %v, want %v", len(g.dealer.Deck), 2)
			}
		})
	}
}
