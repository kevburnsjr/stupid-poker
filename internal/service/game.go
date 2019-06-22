package service

import (
	"strings"
)

const (
	stateReady = iota
	stateDealt

	Hand_RoyalFlush    = "Royal Flush"
	Hand_StraightFlush = "Straight Flush"
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
	// Deal initializes the game, returning the player's initial hand
	Deal() []string

	// Exchange accepts a slice of integers corresponding to indicies for cards in the given hand
	// It discards those cards from the hand and replaces them with new cards from the deck
	// It returns the final hand, a string representation of the hand score and the new balance
	Exchange([]int) ([]string, string, int)

	// GetHandUtf8 returns hand as utf-8 card runes
	GetHandUtf8() []string
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
	if g.state != stateDealt {
		res := scoreHand(g.hand)
		return g.hand, res, g.balance
	}
	hand := g.hand
	for i, n := range cards {
		if n >= len(hand) {
			continue
		}
		hand[n] = g.deck[i+5]
	}
	res := scoreHand(hand)
	g.balance += payout[res]
	g.state = stateReady
	return hand, res, g.balance
}

func (g *game) GetHandUtf8() []string {
	cards := make([]string, 5)
	for i, c := range g.hand {
		card := strings.IndexByte(allCards, c[0])
		suit := string([]rune(c)[1])
		cards[i] = string([]rune(utf8deck[suit])[card])
	}
	return cards
}
