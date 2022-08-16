package pkg

import (
	"encoding/json"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

type basePlayer struct {
	MatchID xid.ID
	RoundID uint
}

func newBasePlayer() *basePlayer {
	p := &basePlayer{}
	p.reset()
	return p
}

func (p *basePlayer) reset() {
	p.MatchID = xid.New()
	p.RoundID = 0
}

func (p *basePlayer) Log(au ArenaUpdate, response string) {
	go p.logSync(au, response)
}

func (p *basePlayer) logSync(au ArenaUpdate, response string) {
	resetMatch := true
	for _, state := range au.Arena.State {
		if state.Score != 0 {
			resetMatch = false
			break
		}
	}

	if resetMatch {
		p.reset()
	}

	_, selfState, err := au.GetSelf()
	if err != nil {
		log.Error().Err(err)
	}

	p.RoundID += 1

	update := Update{
		MatchID:            p.MatchID.String(),
		RoundID:            p.RoundID,
		PreviousRoundScore: uint(selfState.Score),
		Move:               response,
		ArenaUpdate:        au,
	}

	data, err := json.Marshal(update)
	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Msgf("update: %s", data)
}
