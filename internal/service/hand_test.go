package service

import (
	"strings"
	"testing"
)

// Table based test to ensure correct scoring of hands
func TestScoreHand(t *testing.T) {
	var tests = map[string]string{
		"AS KS QS JS TS": Hand_RoyalFlush,
		"JD TD AD QD KD": Hand_RoyalFlush,
		"KS QS JS TS 9S": Hand_StraightFlush,
		"KS KH KD KC 9S": Hand_Quads,
		"KS 9S KH KD KC": Hand_Quads,
		"KS 9S KH KD 9C": Hand_FullHouse,
		"TS 8S QS JS 4S": Hand_Flush,
		"AS KD QS JS TS": Hand_Straight,
		"AS AH AD JS TS": Hand_Trips,
		"AS AH JS AD 4S": Hand_Trips,
		"AS AH JS JD 4S": Hand_TwoPair,
		"AS JD 4S AH JS": Hand_TwoPair,
		"AS JD AH JS 4S": Hand_TwoPair,
		"AS AD 8H JS 4S": Hand_Jacks,
		"JS JD 8H 9S 4S": Hand_Jacks,
		"TS TD 8H 9S 4S": Hand_Nothing,
	}
	for h, exp := range tests {
		res := scoreHand(strings.Split(h, " "))
		if exp != res {
			t.Log("Unexpected Result", h, res, "!=", exp)
			t.Fail()
		}
	}
}
