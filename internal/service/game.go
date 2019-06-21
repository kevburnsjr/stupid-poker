package service

import "log"

const (
	stateReady = iota
	stateDealt

	Hand_RoyalFlush    = "RoyalFlush"
	Hand_StraightFlush = "StraightFlush"
	Hand_Quads         = "Four of a Kind"
	Hand_FullHouse     = "Full House"
	Hand_Flush         = "Flush"
	Hand_Straight      = "Straight"
	Hand_Trips         = "Three of a Kind"
	Hand_TwoPair       = "Two Pair"
	Hand_Jacks         = "Jacks or Higher"
	Hand_Nothing       = "Nothing"
)

var payout = map[string]int{
	Hand_RoyalFlush:    1250,
	Hand_StraightFlush: 250,
	Hand_Quads:         100,
	Hand_FullHouse:     35,
	Hand_Flush:         25,
	Hand_Straight:      20,
	Hand_Trips:         15,
	Hand_TwoPair:       10,
	Hand_Jacks:         5,
	Hand_Nothing:       0,
}

type Game interface {
	Deal() []string
	Exchange([]int) ([]string, string, int)
}

func NewGame(anty int, balance int) Game {
	return &game{
		deck:    newDeck(),
		state:   stateReady,
		anty:    anty,
		balance: balance,
	}
}

type game struct {
	deck    *deck
	state   int
	anty    int
	balance int
	hand    []string
}

func (g *game) Deal() []string {
	g.deck.shuffle()
	g.balance -= g.anty
	g.hand = g.deck[0:5]
	g.state = stateDealt
	return g.hand
}

func (g *game) Exchange(cards []int) ([]string, string, int) {
	hand := g.hand
	for i, n := range cards {
		if n >= len(hand) {
			log.Println("ey")
			continue
		}
		hand[n] = g.deck[i+5]
	}
	res := findResult(hand)
	g.balance += payout[res]
	g.state = stateReady
	g.hand = nil
	return hand, res, g.balance
}
