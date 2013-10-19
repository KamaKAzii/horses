package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Player struct {
	Id    string
	Name  string
	Money int
}

type World struct {
	Players map[string]Player
}

func (world *World) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/api/getPlayers", world.getPlayers)
	mux.ServeHTTP(w, req)
}

func (world *World) getPlayers(w http.ResponseWriter, req *http.Request) {
	ReplyWithJson(w, req, world)
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
	return &World{players}
}

func main() {
	world := NewWorld()
	s := &http.Server{
		Addr: ":8080",
		Handler: world,
	}
	log.Fatal(s.ListenAndServe())
}
