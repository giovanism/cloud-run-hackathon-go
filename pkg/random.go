package pkg

import (
	rand2 "math/rand"
)

type randomPlayer struct {
}

func NewRandomPlayer() Player {
	return &randomPlayer{}
}

func (p *randomPlayer) Play(input ArenaUpdate) (response string) {
	rand := rand2.Intn(4)
	return Moves[rand]
}
