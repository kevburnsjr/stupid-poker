package service

import (
	"math/rand"
	"strings"
	"time"
)

var suits = []string{"♠", "♥", "♦", "♣"}
var cards = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}

var allCards = strings.Join(cards, "")

var utf8deck = map[string]string{
	"♠": "🂡🂢🂣🂤🂥🂦🂧🂨🂩🂪🂫🂭🂮",
	"♥": "🂱🂲🂳🂴🂵🂶🂷🂸🂹🂺🂻🂽🂾",
	"♦": "🃁🃂🃃🃄🃅🃆🃇🃈🃉🃊🃋🃍🃎",
	"♣": "🃑🃒🃓🃔🃕🃖🃗🃘🃙🃚🃛🃝🃞",
}

func newDeck() *deck {
	d := &deck{}
	for i, s := range suits {
		for j, c := range cards {
			d[i*13+j] = c + s
		}
	}
	return d
}

type deck [52]string

// shuffle shuffles the deck in place
func (d *deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(d); n > 0; n-- {
		randIndex := r.Intn(n)
		d[n-1], d[randIndex] = d[randIndex], d[n-1]
	}
}
