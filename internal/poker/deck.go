package poker

import (
	"math/rand"
	"strings"
	"time"
)

var suits = []string{"S", "H", "D", "C"}
var cards = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}

type deck [52]string

func newDeck() *deck {
	d := &deck{}
	for i, s := range suits {
		for j, c := range cards {
			d[i*13+j] = c + s
		}
	}
	return d
}

// shuffles the deck in place
func (d *deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

var allCards = strings.Join(cards, "")

func cardIndex(c byte) int {
	return strings.IndexByte(allCards, c)
}
