package pkg

import (
	rand2 "math/rand"
	"log"
)

type randomPlayer struct {
}

func NewRandomPlayer() Player {
	return &randomPlayer{}
}

func (p *randomPlayer) Play(input ArenaUpdate) (response string) {
	log.Printf("IN: %#v", input)

	rand := rand2.Intn(4)
	return Moves[rand]
}
