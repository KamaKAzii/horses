package main

import (
	"crypto/rand"
	"encoding/binary"
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
	mux.HandleFunc("/api/getWorld", world.getWorld)
	mux.HandleFunc("/api/placeBet", world.placeBet)
	mux.HandleFunc("/api/runRace", world.runRace)
	mux.ServeHTTP(w, req)
}

func (world *World) getWorld(w http.ResponseWriter, req *http.Request) {
	ReplyWithJson(w, req, world)
}

func (world *World) placeBet(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Print(req.Form)
	player_id := req.Form["player"][0]
	horse_id := req.Form["horse"][0]
	amount, err := strconv.Atoi(req.Form["amount"][0])
	if err != nil {
		http.Error(w, "fail", 500)
		return
	}
	world.Bets = append(world.Bets, Bet{player_id, horse_id, amount})
	ReplyWithJson(w, req, "ok")
}

type RunRaceMsg struct {
	WinningHorse string
}

func (world *World) runRace(w http.ResponseWriter, req *http.Request) {
	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	x := int(n % int32(len(world.Horses)))

	ReplyWithJson(w, req, RunRaceMsg{strconv.Itoa(x)})
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
	horses := make(map[string]Horse)
	horses["1"] = Horse{"1", "Crazy Glue"}
	horses["2"] = Horse{"2", "Sickballs"}
	horses["3"] = Horse{"3", "Best Horse"}
	horses["4"] = Horse{"4", "Mr Ed"}
	return &World{players, horses, make([]Bet, 0)}
}

func main() {
	world := NewWorld()
	s := &http.Server{
		Addr:    ":8080",
		Handler: world,
	}
	log.Fatal(s.ListenAndServe())
}
