package poker

import (
	"errors"
	"sync"
)

const (
	stateReady = iota
	stateDealt
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

var ErrNoBalance = errors.New("You are broke")

type Game interface {
	// Deal initializes the game, returning the player's initial hand
	Deal() (Hand, int, error)

	// Exchange accepts a slice of integers corresponding to indicies for cards in the given hand
	// It discards those cards from the hand and replaces them with new cards from the deck
	// It returns the final hand and the new balance
	Exchange([]int) (Hand, int)
}

func NewGame(anty int, balance int) Game {
	return &game{
		deck:    newDeck(),
		state:   stateReady,
		anty:    anty,
		balance: balance,
		mutex:   &sync.Mutex{},
	}
}

type game struct {
	deck    *deck
	state   int
	anty    int
	balance int
	hand    Hand
	mutex   *sync.Mutex
}

func (g *game) Deal() (Hand, int, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.balance < g.anty {
		return Hand{}, g.balance, ErrNoBalance
	}
	g.deck.shuffle()
	g.balance -= g.anty
	g.hand = newHand(g.deck[0:5])
	g.state = stateDealt
	return g.hand, g.balance, nil
}

func (g *game) Exchange(cards []int) (Hand, int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if g.state != stateDealt {
		return g.hand, g.balance
	}
	hand := g.hand
	for i, n := range cards {
		if n >= len(hand) {
			continue
		}
		hand[n] = g.deck[i+5]
	}
	res := hand.Score()
	g.balance += payout[res]
	if payout[res] > 0 {
		g.balance += g.anty
	}
	g.state = stateReady
	return hand, g.balance
}
