package service

import (
	"sort"
	"strings"
)

func findResult(hand []string) string {
	if len(hand) != 5 {
		return Hand_Nothing
	}
	h := append([]string(nil), hand...)
	sort.Sort(byCard(h))
	straight := isStraight(h)
	if isFlush(h) {
		if straight {
			if isRoyal(h) {
				return Hand_RoyalFlush
			}
			return Hand_StraightFlush
		}
		return Hand_Flush
	}
	if straight {
		return Hand_Straight
	}
	pairs, trips, quads, highPair := detectDupes(h)
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

func isRoyal(hand []string) bool {
	for _, c := range hand {
		idx := strings.IndexByte(allCards, c[0])
		if idx > 0 && idx < 9 {
			return false
		}
	}
	return true
}

func isStraight(hand []string) bool {
	var last int
	for i, c := range hand {
		idx := strings.IndexByte(allCards, c[0])
		if i != 0 && idx != last+1 && !(idx == 9 && last == 0) {
			return false
		}
		last = idx
	}
	return true
}

func isFlush(hand []string) bool {
	var suit = []rune(hand[0])[1]
	for _, c := range hand {
		if []rune(c)[1] != suit {
			return false
		}
	}
	return true
}

func detectDupes(hand []string) (int, int, int, bool) {
	var counts = map[byte]int{}
	var highPair bool
	for _, c := range hand {
		counts[c[0]]++
	}
	var pairs int
	var trips int
	var quads int
	for c, n := range counts {
		switch n {
		case 2:
			idx := strings.IndexByte(allCards, c)
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
    return strings.IndexByte(allCards, s[i][0]) < strings.IndexByte(allCards, s[j][0])
}
