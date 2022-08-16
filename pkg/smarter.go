package pkg

import (
	"math"

	"github.com/rs/zerolog/log"
)

type smarterPlayer struct {
	*basePlayer
}

func NewSmarterPlayer() Player {
	return &smarterPlayer{
		basePlayer: newBasePlayer(),
	}
}


// Smarter algorithm inspired by Asing1001
// Ref: https://github.com/Asing1001/cloud-run-hackathon-nodejs
func (p *smarterPlayer) Play(input ArenaUpdate) (response string) {

	defer func(){ p.basePlayer.Log(input, response) }()

	selfUrl, selfState, err := input.GetSelf()
	if err != nil {
		log.Error().Err(err).Msg("failed to get self from input")
		return MoveThrow
	}

	arenaXLength := input.Arena.Dimensions[0]
	arenaYLength := input.Arena.Dimensions[1]

	maxX := arenaXLength - 1
	maxY := arenaYLength - 1

	otherStates := input.Arena.State
	delete(otherStates, selfUrl)

	if selfState.Y == 0 && selfState.Direction == DirectionNorth {
		if selfState.X == 0 {
			return MoveTurnRight
		}

		return MoveTurnLeft
	}

	if selfState.Y == maxY && selfState.Direction == DirectionSouth {
		if selfState.X == 0 {
			return MoveTurnLeft
		}
		
		return MoveTurnRight
	}

	if selfState.X == 0 && selfState.Direction == DirectionWest {
		return MoveTurnLeft
	}

	if selfState.X == maxX && selfState.Direction == DirectionEast {
		return MoveTurnLeft
	}

	throwDistance := 3
	nearStates := make(map[string]PlayerState)

	for url, state := range otherStates {
		if (math.Abs(float64(state.X - selfState.X)) <= float64(throwDistance) && state.Y == selfState.Y) || 
			(math.Abs(float64(state.Y - selfState.Y)) <= float64(throwDistance) && state.X == selfState.X) {
				nearStates[url] = state
			}
	}

	canThrow := false
	for _, state := range nearStates {
		switch selfState.Direction {
			case DirectionEast:
				canThrow = state.Y == selfState.Y && state.X > selfState.X
				continue
			case DirectionWest:
				canThrow = state.Y == selfState.Y && state.X < selfState.X
				continue
			case DirectionNorth:
				canThrow = state.X == selfState.X && selfState.Y > state.Y
				continue
			case DirectionSouth:
				canThrow = state.X == selfState.X && selfState.Y < state.Y
				continue
		}
	}

	moveConflict := func() bool {
		newX := selfState.X
		newY := selfState.Y
		switch selfState.Direction {
			case DirectionEast:
				newX += 1
			case DirectionWest:
				newX -= 1
			case DirectionNorth:
				newY += 1
			case DirectionSouth:
				newY -= 1
		}

		conflict := false
		for _, state := range nearStates {
			conflict = state.X == newX && state.Y == newY
		}

		return conflict
	}

	if canThrow {
		if selfState.WasHit {
			if moveConflict() {
				return MoveTurnLeft
			}

			return MoveForward
		}

		return MoveThrow
	} else {
		if len(nearStates) > 0 && !selfState.WasHit {
			return MoveTurnLeft
		}
		if moveConflict() {
			return MoveTurnLeft
		}

		return MoveForward
	}
}
