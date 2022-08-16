package pkg

import (
	rand2 "math/rand"
)

type randomPlayer struct {
	*basePlayer
}

func NewRandomPlayer() Player {
	return &randomPlayer{
		basePlayer: newBasePlayer(),
	}
}

func (p *randomPlayer) Play(input ArenaUpdate) (response string) {
	defer p.basePlayer.Log(input)
	rand := rand2.Intn(4)
	return Moves[rand]
}
