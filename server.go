package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type WorldServer struct {
	world *World
}

func (s *WorldServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/api/getWorld", s.getWorld)
	mux.HandleFunc("/api/placeBet", s.placeBet)
	mux.HandleFunc("/api/runRace", s.runRace)
	mux.ServeHTTP(w, req)
}

func (s *WorldServer) getWorld(w http.ResponseWriter, req *http.Request) {
	replyWithJson(w, req, s.world)
}

func (s *WorldServer) placeBet(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Print(req.Form)
	playerId := req.Form["player"][0]
	horseId := req.Form["horse"][0]
	amount, err := strconv.Atoi(req.Form["amount"][0])
	if err != nil {
		http.Error(w, "Enter an amount before placing a bet.", 500)
		return
	}
	s.world.placeBet(playerId, horseId, amount)
	replyWithJson(w, req, "ok")
}

type RunRaceMsg struct {
	RaceOrder []string
}

func (s *WorldServer) runRace(w http.ResponseWriter, req *http.Request) {
	order := rand.Perm(len(s.world.Horses))
	horsePositions := make([]string, len(s.world.Horses))
	horseIds := make([]string, len(s.world.Horses))
	i := 0
	for _, h := range s.world.Horses {
		horseIds[i] = h.Id
		i++
	}
	for i, n := range order {
		horsePositions[i] = horseIds[n]
	}
	s.world.processWinner(horsePositions[0])
	replyWithJson(w, req, RunRaceMsg{horsePositions})
}

func replyWithJson(w http.ResponseWriter, req *http.Request, i interface{}) {
	w.WriteHeader(200)
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "fail", 500)
		return
	}
	w.Write(data)
}

func NewWorld() (world *World) {
	players := make(map[string]*Player)
	players["1"] = &Player{"1", "James", 100}
	players["2"] = &Player{"2", "Emmet", 100}
	players["3"] = &Player{"3", "Michelle", 100}
	horses := make(map[string]*Horse)
	horses["1"] = &Horse{"1", "Crazy Glue"}
	horses["2"] = &Horse{"2", "Sickballs"}
	horses["3"] = &Horse{"3", "Best Horse"}
	horses["4"] = &Horse{"4", "Mr Ed"}
	return &World{players, horses, make([]Bet, 0)}
}

func main() {
	world := NewWorld()
	s := &http.Server{
		Addr:    ":8080",
		Handler: &WorldServer{world},
	}
	fmt.Println("Listening on :8080")
	log.Fatal(s.ListenAndServe())
}
