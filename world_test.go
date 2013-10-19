package main

import (
	"fmt"
	"testing"
)

func makeTestWorld() *World {
	players := make(map[string]*Player)
	players["1"] = &Player{"1", "James", 100}
	players["2"] = &Player{"2", "Emmet", 100}
	horses := make(map[string]*Horse)
	horses["1"] = &Horse{"1", "Crazy Horse"}

	return &World{
		Players: players,
		Horses: horses,
		Bets: make([]Bet, 0),
	}
}

func TestPlaceBet(t *testing.T) {
	world := makeTestWorld()
	world.placeBet("1", "1", 20)
	if world.Players["1"].Money != 80 {
		t.Fail()
	}
	if world.Players["2"].Money != 100 {
		t.Fail()
	}
	if len(world.Bets) != 1 {
		t.Fail()
	}
	if world.Bets[0].HorseId != "1" {
		t.Fail()
	}
}

func TestProcessWinner(t *testing.T) {
	world := makeTestWorld()
	world.placeBet("1", "1", 20)
	world.processWinner("1")
	if len(world.Bets) > 0 {
		t.Fail()
	}
	fmt.Printf("%d\n", world.Players["1"].Money)
	if world.Players["1"].Money != 120 {
		t.Fail()
	}
}
