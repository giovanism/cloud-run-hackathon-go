package pkg

import (
	"fmt"
)

type ArenaUpdate struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Arena struct {
		Dimensions []int                  `json:"dims"`
		State      map[string]PlayerState `json:"state"`
	} `json:"arena"`
}

// GetSelf url and state
func (au *ArenaUpdate) GetSelf() (string, PlayerState, error) {
	url := au.Links.Self.Href
	state, ok := au.Arena.State[url]
	if !ok {
		err := fmt.Errorf("self state not found")
		return url, PlayerState{}, err
	}

	return url, state, nil
}

type PlayerState struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Direction string `json:"direction"`
	WasHit    bool   `json:"wasHit"`
	Score     int    `json:"score"`
}

type Player interface {
	Play(ArenaUpdate) string
}

const (
	DirectionNorth = "N"
	DirectionWest  = "W"
	DirectionSouth = "S"
	DirectionEast  = "E"

	MoveForward   = "F"
	MoveThrow     = "T"
	MoveTurnLeft  = "L"
	MoveTurnRight = "R"
)

var (
	Directions = []string{DirectionNorth, DirectionWest, DirectionSouth, DirectionEast}
	Moves      = []string{MoveForward, MoveThrow, MoveTurnLeft, MoveTurnRight}
)

type Update struct {
	MatchID            string      `json:"match_id"`
	RoundID            uint        `json:"round_id"`
	PreviousRoundScore int         `json:"previous_round_score"`
	Move               string      `json:"move"`
	ArenaUpdate        ArenaUpdate `json:"arena_update"`
}
