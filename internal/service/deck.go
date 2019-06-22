package service

import (
	"math/rand"
	"strings"
	"time"
)

var suits = []string{"S", "H", "D", "C"}
var cards = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}

var allCards = strings.Join(cards, "")

var utf8deck = map[string]string{
	"S": "🂡🂢🂣🂤🂥🂦🂧🂨🂩🂪🂫🂭🂮",
	"H": "🂱🂲🂳🂴🂵🂶🂷🂸🂹🂺🂻🂽🂾",
	"D": "🃁🃂🃃🃄🃅🃆🃇🃈🃉🃊🃋🃍🃎",
	"C": "🃑🃒🃓🃔🃕🃖🃗🃘🃙🃚🃛🃝🃞",
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
	for n := 0; n < len(d); n++ {
		randIndex := r.Intn(len(d))
		d[n], d[randIndex] = d[randIndex], d[n]
	}
}
