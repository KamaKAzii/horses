package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Player struct {
	Id    string
	Name  string
	Money int
}

type Horse struct {
	Id   string
	Name string
}

type Bet struct {
	PlayerId string
	HorseId  string
	Amount   int
}

type World struct {
	Players map[string]Player
	Horses  map[string]Horse
	Bets    []Bet
}

func (world *World) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/api/getPlayers", world.getPlayers)
	mux.HandleFunc("/api/placeBet", world.placeBet)
	mux.ServeHTTP(w, req)
}

func (world *World) getPlayers(w http.ResponseWriter, req *http.Request) {
	ReplyWithJson(w, req, world)
}

func (world *World) placeBet(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Print(req.Form)
	player_id := req.Form["Player"][0]
	horse_id := req.Form["Horse"][0]
	amount, err := strconv.Atoi(req.Form["Amount"][0])
	if err != nil {
		http.Error(w, "fail", 500)
		return
	}
	world.Bets = append(world.Bets, Bet{player_id, horse_id, amount})
	ReplyWithJson(w, req, "ok")
}

func ReplyWithJson(w http.ResponseWriter, req *http.Request, i interface{}) {
	w.WriteHeader(200)
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "fail", 500)
		return
	}
	w.Write(data)
}

func NewWorld() (world *World) {
	players := make(map[string]Player)
	players["1"] = Player{"1", "James", 100}
	players["2"] = Player{"2", "Emmet", 100}
	players["3"] = Player{"3", "Michelle", 100}
	return &World{players, make(map[string]Horse), make([]Bet, 0)}
}

func main() {
	world := NewWorld()
	s := &http.Server{
		Addr:    ":8080",
		Handler: world,
	}
	log.Fatal(s.ListenAndServe())
}
