package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

type basePlayer struct {
	MatchID xid.ID
	RoundID uint
	Score   int
}

func newBasePlayer() *basePlayer {
	p := &basePlayer{}
	p.reset()
	return p
}

func (p *basePlayer) reset() {
	log.Info().
		Str("match_id", p.MatchID.String()).
		Uint("round_id", p.RoundID).
		Int("score", p.Score).
		Msg("match final result")

	p.MatchID = xid.New()
	p.RoundID = 0
	p.Score = 0
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
		log.Error().Err(err).Msg("failed to get self state")
	}

	p.RoundID += 1
	p.Score = selfState.Score

	update := Update{
		MatchID:            p.MatchID.String(),
		RoundID:            p.RoundID,
		PreviousRoundScore: selfState.Score,
		Move:               response,
		ArenaUpdate:        au,
	}

	data, err := json.Marshal(update)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal update data")
	}

	log.Info().
		Str("update", fmt.Sprintf("%s", data)).
		Msg("update response")
}
