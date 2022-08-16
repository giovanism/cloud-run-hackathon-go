package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	play "github.com/giovanism/cloud-run-hackathon-go/pkg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	if env := os.Getenv("ENV"); env == "" || env == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	// else env == "prod" will use default/json writer
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
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read request body")
	}

	log.Info().
		Str("arena_update", fmt.Sprintf("%s", data)).
		Msg("raw request body")

	if err := json.Unmarshal(data, &v); err != nil {
		log.Warn().Err(err).Msg("failed to unmarshal ArenaUpdate in request body data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	resp := player.Play(v)
	fmt.Fprint(w, resp)
}
