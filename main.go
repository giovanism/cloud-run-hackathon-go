package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	play "github.com/giovanism/cloud-run-hackathon-go/pkg"
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
		log.Error().Err(err)
	}

	log.Info().Msgf("raw: %s", data)

	if err := json.Unmarshal(data, &v); err != nil {
		log.Warn().Err(err).Msg("failed to unmarshal ArenaUpdate in response body data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	resp := player.Play(v)
	fmt.Fprint(w, resp)
}
