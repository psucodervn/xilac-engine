package xilac

import "errors"

var (
	ErrInvalidPlayerCount = errors.New("require at least 2 players")
)

type GameState struct {
	Deck             Deck `json:"deck"`
	CurrentPlayerIdx int  `json:"current_player_idx"`
}

type GamePlayerStatus uint8

const (
	GamePlayerStatusWaiting GamePlayerStatus = iota
	GamePlayerStatusPlaying
	GamePlayerStatusFinished

	GamePlayerStatusDoubleBlackJack = iota + 97 // NOTE: ensure this is equal to 100
	GamePlayerStatusBlackJack
	GamePlayerStatusHighFive
	GamePlayerStatusNormal
	GamePlayerStatusBusted
	GamePlayerStatusTooHigh
	GamePlayerStatusTooLow
)

type GamePlayer struct {
	Deck   Deck             `json:"deck"`
	Status GamePlayerStatus `json:"status"`
}

type Game struct {
	deck    Deck
	players []GamePlayer
	dealer  GamePlayer
}

func NewGame(playerCount int) (*Game, error) {
	if playerCount < 2 {
		return nil, ErrInvalidPlayerCount
	}

	players := make([]GamePlayer, playerCount)
	dealer := GamePlayer{}

	return &Game{
		players: players,
		dealer:  dealer,
	}, nil
}

func (g *Game) deal() {
	if len(g.deck) > 0 {
		panic("already dealt")
	}

	g.deck = NewShuffledFullDeck()
	for step := 1; step <= 2; step++ {
		for i := range g.players {
			var card Card
			card, g.deck = g.deck.Pop()
			g.players[i].Deck = g.players[i].Deck.Push(card)
		}
		var card Card
		card, g.deck = g.deck.Pop()
		g.dealer.Deck = g.dealer.Deck.Push(card)
	}
}
