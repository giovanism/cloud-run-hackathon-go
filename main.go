package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
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
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := player.Play(v)
	fmt.Fprint(w, resp)
}
