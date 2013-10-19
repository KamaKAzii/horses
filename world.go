package main

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
	Players map[string]*Player
	Horses  map[string]*Horse
	Bets    []Bet
}

func (world *World) processWinner(winnerId string) {
	bets := world.Bets
	world.Bets = make([]Bet, 0)
	winFactor := len(world.Horses)
	for _, b := range bets {
		if b.HorseId == winnerId {
			// TODO(koz): Implement odds.
			world.Players[b.PlayerId].Money += winFactor * b.Amount
		}
	}
}

func (world *World) placeBet(playerId string, horseId string, amount int) {
	world.Players[playerId].Money -= amount
	world.Bets = append(world.Bets, Bet{playerId, horseId, amount})
}
