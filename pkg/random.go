package pkg

import (
	rand2 "math/rand"
	"log"
)

func Random(input ArenaUpdate) (response string) {
	log.Printf("IN: %#v", input)

	rand := rand2.Intn(4)
	return Moves[rand]
}
