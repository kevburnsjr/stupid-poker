package poker

import (
	"sort"
)

const (
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

type Hand [5]string

func newHand(cards []string) Hand {
	h := Hand{}
	for i, c := range cards {
		if i >= 5 {
			break
		}
		h[i] = c
	}
	return h
}

// Score scores the hand
func (h Hand) Score() string {
	if len(h[0]) < 1 {
		return Hand_Nothing
	}
	if h.isFlush() {
		if h.isRoyal() {
			return Hand_RoyalFlush
		}
		if h.isStraight() {
			return Hand_StraightFlush
		}
		return Hand_Flush
	}
	if h.isStraight() {
		return Hand_Straight
	}
	pairs, trips, quads, highPair := h.detectDupes()
	if quads > 0 {
		return Hand_Quads
	}
	if trips > 0 {
		if pairs > 0 {
			return Hand_FullHouse
		}
		return Hand_Trips
	}
	if pairs > 1 {
		return Hand_TwoPair
	}
	if pairs == 1 && highPair {
		return Hand_Jacks
	}
	return Hand_Nothing
}

func (h Hand) clone() Hand {
	h2 := Hand{}
	for i := range h {
		h2[i] = h[i]
	}
	return h2
}

func (h Hand) isRoyal() bool {
	for _, c := range h {
		idx := cardIndex(c[0])
		if idx > 0 && idx < 9 {
			return false
		}
	}
	return true
}

func (h Hand) isStraight() bool {
	hs := h.clone()
	sort.Sort(byCard(hs[:]))
	var last int
	for i, c := range hs {
		idx := cardIndex(c[0])
		if i != 0 && idx != last+1 && !(idx == 9 && last == 0) {
			return false
		}
		last = idx
	}
	return true
}

func (h Hand) isFlush() bool {
	var suit = []rune(h[0])[1]
	for _, c := range h {
		if []rune(c)[1] != suit {
			return false
		}
	}
	return true
}

func (h Hand) detectDupes() (int, int, int, bool) {
	var counts = map[byte]int{}
	var highPair bool
	for _, c := range h {
		counts[c[0]]++
	}
	var pairs int
	var trips int
	var quads int
	for c, n := range counts {
		switch n {
		case 2:
			idx := cardIndex(c)
			if idx == 0 || idx > 9 {
				highPair = true
			}
			pairs++
		case 3:
			trips++
		case 4:
			quads++
		}
	}
	return pairs, trips, quads, highPair
}

type byCard []string

func (s byCard) Len() int {
	return len(s)
}
func (s byCard) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byCard) Less(i, j int) bool {
	return cardIndex(s[i][0]) < cardIndex(s[j][0])
}
