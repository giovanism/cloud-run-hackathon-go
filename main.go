package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	play "github.com/giovanism/cloud-run-hackathon-go/pkg"
)

var player play.Player

func init() {
	var strategy string
	if v := os.Getenv("STRATEGY"); v != "" {
		 strategy = strings.ToLower(v)
	}

	switch strategy {
	case "smarter":
		player = play.NewSmarterPlayer()
	default:
		player = play.NewRandomPlayer()
	}
}

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)

	log.Info().Msgf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal().Msgf("http listen error: %v", err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprint(w, "Let the battle begin!")
		return
	}

	var v play.ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Warn().Err(err).Msg("failed to decode ArenaUpdate in response body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info().Msgf("arena update", v)
	resp := player.Play(v)
	fmt.Fprint(w, resp)
}
